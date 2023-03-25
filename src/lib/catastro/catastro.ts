import { XMLParser } from 'fast-xml-parser';

class Coor {
	x: string;
	y: string;
	constructor(x: string, y: string) {
		this.x = x;
		this.y = y;
	}
}

class Subparcelas {
	calificacion: string;
	superficie: number;
	intensidadProductiva: number;
	tipoUso: string;
	constructor(
		calificacion: string,
		superficie: number,
		intensidadProductiva: number,
		tipoUso: string
	) {
		this.calificacion = calificacion;
		this.superficie = superficie;
		this.intensidadProductiva = intensidadProductiva;
		this.tipoUso = tipoUso;
	}
}

export function validateRC(rc: string): boolean {
	// 13077A01800039
	// Capture groups:
	// Provincia. 2 digits
	// Municipio. 3 digits
	// Sector. 1 letter
	// Polígono. 3 digits
	// Parcela. 5 digits
	const regex = /^(\d{2})(\d{3})([A-Z])(\d{3})(\d{5})$/;
	const match = rc.match(regex);
	if (match) {
		return true;
	}
	return false;
}
export class Catastro {
	rc: string;
	coor: Coor;
	provincia: string;
	municipio: string;
	direccion: string;
	numero: string;
	uso: string;
	subparcelas: Subparcelas[];

	constructor(rc: string) {
		if (!validateRC(rc)) {
			throw Error('RC is not valid');
		}
		this.rc = rc;
		this.coor = new Coor('0', '0');
		this.provincia = '';
		this.municipio = '';
		this.direccion = '';
		this.numero = '';
		this.uso = '';
		this.subparcelas = [];
	}
	// Catastro validation

	async getRemoteData(): Promise<any | Error> {
		const data = await this.getRawData();
		if (data instanceof Error) {
			console.error(data);
			throw Error('Stupid Error');;
		} else {
			await this.mapCatatroData(data);
		}
	}

	async getRawData(): Promise<any | Error> {
		const responseM2 = await fetch(
			'http://ovc.catastro.meh.es/ovcservweb/OVCSWLocalizacionRC/OVCCallejero.asmx/Consulta_DNPRC',
			{
				headers: {
					accept: 'application/xml',
					'cache-control': 'max-age=0',
					'content-type': 'application/x-www-form-urlencoded',
					'upgrade-insecure-requests': '1'
				},
				body: `MUNICIPIO=&RC=${this.rc}&PROVINCIA=`,
				method: 'POST'
			}
		);
		

		if (responseM2.status !== 200) {
			return Error('Error getting data from catastro');
		}

		let XMLData = await responseM2.text();
		const parser = new XMLParser();
		let jObj = parser.parse(XMLData);
		if (XMLData.includes('Error')) {
			return Error('Error getting data from catastro');
		}
		return jObj;
	}

	async mapCatatroData(data: any) {
		this.setMunicipio(data);
		this.setProvincia(data);
		this.setDireccion(data);
		this.setUso(data);
		this.setSubparcelas(data);
		await this.setRemoteCoor();
	}

	setMunicipio(data: any) {
		const bi = data.consulta_dnp.bico.bi;
		const dt = bi.dt;
		const nm = dt.nm;

		this.municipio = nm;
	}

	setProvincia(data: any) {
		const bi = data.consulta_dnp.bico.bi;
		const dt = bi.dt;
		const np = dt.nm;
		this.provincia = np;
	}

	setDireccion(data: any) {
		const bi = data.consulta_dnp.bico.bi;
		this.direccion = bi.ldt;
	}

	setUso(data: any) {
		const bi = data.consulta_dnp.bico.bi;
		const debi = bi.debi;
		const luso = debi.luso;
		this.uso = luso;
	}

	setSubparcelas(data: any) {
		const subpar: Subparcelas[] = [];
		if (Array.isArray(data.consulta_dnp.bico.lspr.spr)) {
			const elements = data.consulta_dnp.bico.lspr.spr;
			elements.forEach((element: any) => {
				subpar.push(
					new Subparcelas(
						element.dspr.ccc as string,
						element.dspr.ssp as number,
						element.dspr.ip as number,
						element.dspr.dcc as string
					)
				);
			});
		} else {
			const element = data.consulta_dnp.bico.lspr.spr.dspr;
			subpar.push(
				new Subparcelas(
					element.ccc as string,
					element.ssp as number,
					element.ip as number,
					element.dcc as string
				)
			);
		}

		this.subparcelas = subpar;
	}

	// Get total m2 of all subparcelas
	getTotalSuperficie() {
		return this.subparcelas.reduce((acc: number, item: Subparcelas) => {
			return acc + item.superficie;
		}, 0);
	}

	async setRemoteCoor(): Promise<any | Error> {
		const url = `http://ovc.catastro.meh.es/OVCServWeb/OVCWcfCallejero/COVCCoordenadas.svc/rest/Consulta_CPMRC?SRS=EPSG:4326&RefCat=${this.rc}`;
		const responseCoor = await fetch(url);
		if (responseCoor.status !== 200) {
			return Error('Error getting data from catastro');
		}
		let XMLData = await responseCoor.text();
		const parser = new XMLParser();
		const JSONData = parser.parse(XMLData);
		const x = JSONData.consulta_coordenadas.coordenadas.coord.geo.xcen;
		const y = JSONData.consulta_coordenadas.coordenadas.coord.geo.ycen;
		this.coor = new Coor(x, y);
		return this.coor;
	}
}

export async function getGeoDataFromRC(rc: string) {
	// const cat = new Catastro(rc);
	// await cat.getRemoteData();
	// https://ovc.catastro.meh.es/INSPIRE/wfsCP.aspx?service=wfs&version=2&request=getfeature&STOREDQUERIE_ID=GetParcel&refcat=33034A02000032&srsname=EPSG::4326
	const url = `http://ovc.catastro.meh.es/INSPIRE/wfsCP.aspx?service=wfs&version=2&request=getfeature&STOREDQUERIE_ID=GetParcel&refcat=${rc}&srsname=EPSG::4326`;
	const response = await fetch(url);

	if (response.status !== 200) {
		return Error('Error getting data from catastro');
	}
	let XMLData = await response.text();
	const parser = new XMLParser();
	const JSONData = parser.parse(XMLData);
	// {"?xml":"","FeatureCollection":{"member":{"cp:CadastralParcel":{"cp:areaValue":1806,"cp:beginLifespanVersion":"2014-10-08T00:00:00","cp:endLifespanVersion":"","cp:geometry":{"gml:MultiSurface":{"gml:surfaceMember":{"gml:Surface":{"gml:patches":{"gml:PolygonPatch":{"gml:exterior":{"gml:LinearRing":{"gml:posList":"43.545363 -6.549606 43.545425 -6.549432 43.545548 -6.549089 43.545533 -6.549083 43.545458 -6.549088 43.545345 -6.549081 43.545244 -6.549071 43.545132 -6.549065 43.545059 -6.549065 43.545003 -6.549072 43.544969 -6.549082 43.544911 -6.549291 43.545363 -6.549606"}},"gml:interior":{"gml:LinearRing":{"gml:posList":"43.545183 -6.549235 43.545145 -6.549336 43.545096 -6.549302 43.545092 -6.549313 43.545063 -6.549292 43.545087 -6.549225 43.545073 -6.549215 43.54509 -6.54917 43.545183 -6.549235"}}}}}}}},"cp:inspireId":{"Identifier":{"localId":"33034A02000032","namespace":"ES.SDGC.CP"}},"cp:label":32,"cp:nationalCadastralReference":"33034A02000032","cp:referencePoint":{"gml:Point":{"gml:pos":"43.54524 -6.549273"}}}}}}
	const polygonRaw =
		JSONData.FeatureCollection.member['cp:CadastralParcel']['cp:geometry']['gml:MultiSurface'][
			'gml:surfaceMember'
		]['gml:Surface']['gml:patches']['gml:PolygonPatch']['gml:exterior']['gml:LinearRing'][
			'gml:posList'
		];
	const coords = polygonRaw.split(' ');
	// Map array to array of arrays [ [x,y], [x,y], [x,y]
	const coordsArray: number[][] = coords.reduce((acc: number[][], item: number, index: number) => {
		if (index % 2 === 0) {
			acc.push([item]);
		} else {
			acc[acc.length - 1].push(item);
		}
		return acc;
	}, []);

	console.log(JSON.stringify(JSONData));
	console.log(JSON.stringify(coordsArray));
	return coordsArray;
}
