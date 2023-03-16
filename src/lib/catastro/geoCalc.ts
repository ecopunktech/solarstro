import { area, squareGrid, bbox, type BBox, type Properties,polygon,booleanContains } from '@turf/turf';
import {getGeoDataFromRC} from './catastro';
export function getArea(pol: number[][][]): number {
    // getGeoDataFromRC('33034A02000032');
	const outerPolygon = polygon(pol)
    console.log(outerPolygon.geometry.coordinates);
    const area2 = area(outerPolygon);
    console.log(area2);
	const bboxOut: BBox = bbox(outerPolygon);
	const cellSide = 0.01;
	const options: Properties = { units: 'kilometers', mask: outerPolygon };
	const squareGrids = squareGrid(bboxOut, cellSide);
    console.log(outerPolygon.geometry.coordinates);
    const insideResult = squareGrids.features.filter((feature) => {
        console.log(booleanContains(outerPolygon, feature));
        console.log(feature.geometry.coordinates);
        return booleanContains(outerPolygon, feature);
    });
    return insideResult.length;
}