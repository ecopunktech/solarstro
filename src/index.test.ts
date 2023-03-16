import { describe, it, expect } from 'vitest';
import { getArea } from '$lib/catastro/geoCalc';
import { getGeoDataFromRC } from '$lib/catastro/catastro';
describe('sum test', () => {
	it('adds 1 + 2 to equal 3', () => {
		expect(1 + 2).toBe(3);
	});
	it('get proper grid', () => {
		const pol = [
			[
				[0, 0],
				[0, 1],
				[1, 1],
				[1, 0],
				[0, 0]
			]
		];
		// const grid = getArea(pol);
		// expect(getArea(pol)).toBe(1);
	});
	it('get proper grid', async () => {
		const coor = await getGeoDataFromRC('33034A02000030');
		console.log(coor);
		if (coor instanceof Error) {
			console.log(coor);
		} else {
		const pannel = getArea([coor]);
		console.log(pannel);
		expect(pannel).toBe(1);
		}
	});
});
