import { parseStringPromise } from 'xml2js';

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
		// return result;
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
				method: 'POST',
				mode: 'cors',
				credentials: 'include'
			}
		);
		if (responseM2.status !== 200) {
			console.log(responseM2.status);
			return Error('Error getting data from catastro');
		}
		let XMLData = await responseM2.text();
		if (XMLData.includes('Error')) {
			return Error('Error getting data from catastro');
		}
		return parseStringPromise(XMLData);
	}

	async mapCatatroData(data: any) {
		const nParcelas = data.consulta_dnp.control.reduce((acc: any, item: any) => {
			return item.cudnp;
		});

		const nSubparcelas = data.consulta_dnp.control.reduce((acc: any, item: any) => {
			return item.cucul;
		});
		this.setMunicipio(data);
		this.setProvincia(data);
		this.setDireccion(data);
		this.setUso(data);
		this.setSubparcelas(data);
		await this.setRemoteCoor();
	}

	setMunicipio(data: any) {
		const bi = data.consulta_dnp.bico.reduce((acc: any, item: any) => {
			return item.bi;
		});
		const dt = bi.bi.reduce((acc: any, item: any) => {
			return item.dt;
		});
		const nm = dt.dt.reduce((acc: any, item: any) => {
			return item.nm;
		});

		this.municipio = nm.nm.shift() as string;
	}

	setProvincia(data: any) {
		const bi = data.consulta_dnp.bico.reduce((acc: any, item: any) => {
			return item.bi;
		});
		const dt = bi.bi.reduce((acc: any, item: any) => {
			return item.dt;
		});
		const np = dt.dt.reduce((acc: any, item: any) => {
			return item.np;
		});
		this.provincia = np.np.shift() as string;
	}

	setDireccion(data: any) {
		// this.direccion = data.consulta_dnp.bico.reduce((acc: any, item: any) => {
		// 	return item.bi.reduce((acc: any, item: any) => {
		// 		return item.ldt.shift() as string;
		// 	});
		// });

		const bi = data.consulta_dnp.bico.reduce((acc: any, item: any) => {
			return item.bi;
		});
		const ldt = bi.bi.reduce((acc: any, item: any) => {
			return item.ldt;
		});
		this.direccion = ldt.ldt.shift() as string;
	}

	setUso(data: any) {
		const bi = data.consulta_dnp.bico.reduce((acc: any, item: any) => {
			return item.bi;
		});
		const debi = bi.bi.reduce((acc: any, item: any) => {
			return item.debi;
		});
		const luso = debi.debi.reduce((acc: any, item: any) => {
			return item.luso;
		});
		this.uso = luso.luso.shift() as string;
	}

	setSubparcelas(data: any) {
		const subpar: Subparcelas[] = [];
		data.consulta_dnp.bico.forEach((element: any) => {
			element.lspr.forEach((element: any) => {
				element.spr.forEach((element: any) => {
					element.dspr.forEach((element: any) => {
						subpar.push(
							new Subparcelas(
								element.ccc.shift() as string,
								element.ssp.shift() as number,
								element.ip.shift() as number,
								element.dcc.shift() as string
							)
						);
					});
				});
			});
		});
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
		const XMLData = await responseCoor.text();
		if (XMLData.includes('Error')) {
			return Error('Error getting data from catastro');
		}
		const JSONData = await parseStringPromise(XMLData);
		// TODO: Check if we need to get all coordinates or just the first one
		// console.log(JSON.stringify(JSONData));
		const coordinates = JSONData.consulta_coordenadas.coordenadas.reduce((acc: any, item: any) => {
			return item.geo;
		});
		const geo = coordinates.coord.reduce((acc: any, item: any) => {
			return new Coor(item.xcen.shift(), item.ycen.shift());
		});
		const coor = geo.geo.reduce((acc: any, item: any) => {
			return new Coor(item.xcen.shift(), item.ycen.shift());
		});
		const x = coor.xcen.shift();
		const y = coor.ycen.shift();
		this.coor = new Coor(x, y);
		return this.coor;
		// const lon = JSONData.coordenadas.X;
		// this.coor = new Coor(JSONData.Coordenadas.X, JSONData.Coordenadas.Y);
	}
}