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
		this.rc = rc;
		this.coor = new Coor('0', '0');
		this.provincia = '';
		this.municipio = '';
		this.direccion = '';
		this.numero = '';
		this.uso = '';
		this.subparcelas = [];
	}

	async getRemoteData() {
		const data = await this.getRawData();
		if (data instanceof Error) {
			console.log(data);
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
		const element = data.consulta_dnp.bico.lspr.spr.dspr;
		subpar.push(
			new Subparcelas(
				element.ccc as string,
				element.ssp as number,
				element.ip as number,
				element.dcc as string
			)
		);

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
