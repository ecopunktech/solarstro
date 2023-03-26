# solarstro

Este repositorio contiene el código de un backend en Go para calcular una granja solar. Se utiliza el registro catastral y se conecta a https://re.jrc.ec.europa.eu/pvg_tools/en/ para obtener información sobre la radiación solar.

## Requerimientos

 - Go 1.16 o superior
 - Acceso a Internet para conectar con https://re.jrc.ec.europa.eu/

## Instalación

1. Clona el repositorio en tu máquina local.
```
git clone https://github.com/ecopunktech/solarstro.git
```
2. Navega hasta la carpeta `backend`.
```
cd solarstro/backend
```
3. Ejecuta el comando `go run main.go` para iniciar el servidor en `http://localhost:8080`.

## Uso

El servidor tiene dos endpoints:
- `/report`: Retorna un objeto JSON con la respuesta del cálculo.
- `/svg`: Retorna un archivo SVG con la representación de la granja solar.

El código utiliza los siguientes paneles solares para el cálculo:

| Nº | Fabricante  | Modelo                   | Tipo             | Eficiencia | Potencia en Watios | Tensión de Conexión | Tamaño en mm    | Precio en €. Iva incluido | Garantía en años | Observaciones |
|----|-------------|--------------------------|-------------------|------------|--------------------------|-------------------------|----------------|---------------------------------|------------------|----------------|
| 1  | Sun Power   | Max3 400 Wp              | Monocristalino | Hasta 22%   | 400                          | 24 Voltios                    | 1046x1690x40 | 375                             | 40/40               | 72 celdas            |
| 2  | Panasonic (Sanyo) | VBHN325SJ47 | Monocristalino | Hasta 19%   | 325                          | 24/32 Voltios                  | 1053x1590x35 | 250                             | 25/10               | 96 celdas            |
| 3  | LG          | Neon R LG370Q1C-V5 | Monocristalino | "Hasta 17%" | 370                          | 12 Voltios                    | 1016x1700x40 | 300                             | 25/10               | 60 celdas            |
| 4  | Sharp       | NU-AH370                | Monocristalino | Hasta 19%   | 370                          | 24 Voltios                    | 1956x992x35   | 160                             | 25/10               | 72 celdas            |
| 5  | JinkoSolar  | Tiger Pro 460Wp 120cél 1500v  | Monocristalino | Hasta 18,7% | 460                          | 24 Voltios                    | 1903x1134x30 | 160 | 25/12              | 120 celdas          |
| 6  | RisenEnergy | Perc Titan 40            | Monocristalino | Hasta 21,6% | 450                          | 24 Voltios                     | 1754x1096x30 | 178                             | 25/12              | 120 celdas          |
| 7  | EGingPV     | Star Pro EG-460NT60-HLV | Monocristalino | Hasta 21,24% | 460                         | 24 Voltios                     | 1909x1134x30 | 180                             | 25/12              | 120 celdas          |
| 8  | Atersa      | A-450M GS (M6x24) PERC  | Monocristalino | Hasta 20,7% | 450                          | 24 Voltios                    | 2094 × 1038x40 | 192                            | 25/10              | 72 celdas           |
| 9  | Eurener     | MEPV144.HALF-CUT PLUS   | Monocristalino | Hasta 21,3% | 450                          | 24 Voltios                    | 2094x1038x35  | 160                             | 25/20              | 72 celdas           |
| 10 | EastchSolar | Amerisolar 450          | Monocristalino | Hasta 20,58% | 450                        | 24 Voltios                     | 2102x1040x35 | 295                             | 25/12              | 72 celdas           |
| 11 | Tamesol     | TM 430-450M-144HC

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).

## Autor

Este proyecto fue desarrollado por [@ecopunktech](https://github.com/ecopunktech).