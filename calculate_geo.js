// 697985.83 4824277.34 697991.33 4824267.19 697987.76 4824265.25 697986.92 4824266.79 697981.56 4824263.88 697979.84 4824267.04 697980.72 4824267.52 697977.77 4824272.97 697985.83 4824277.34

const stringPoints =
	'697955.28 4824296.53 697969.11 4824303.83 697996.44 4824318.27 697996.94 4824316.65 697996.78 4824308.28 697997.76 4824295.69 697998.85 4824284.52 697999.74 4824272.09 697999.93 4824264.02 697999.59 4824257.78 697998.9 4824253.98 697982.16 4824247.09 697955.28 4824296.53';

const points = stringPoints
	.split(' ')
	.map((point, index) => {
		if (index % 2 === 0) {
			return [point, stringPoints.split(' ')[index + 1]];
		}
	})
	.filter((point) => point !== undefined);
const points2 = [
	[0, 0],
	[10, 0],
	[10, 10],
	[0, 10]
];
const distance = (point1, point2) => {
	const x = point2[0] - point1[0];
	const y = point2[1] - point1[1];
	return Math.sqrt(x * x + y * y);
};

function calculateBoundingBox(polygon) {
	let minX = Infinity;
	let maxX = -Infinity;
	let minY = Infinity;
	let maxY = -Infinity;

	for (const point of polygon) {
		minX = Math.min(minX, point[0]);
		maxX = Math.max(maxX, point[0]);
		minY = Math.min(minY, point[1]);
		maxY = Math.max(maxY, point[1]);
	}

	return { minX, maxX, minY, maxY };
}

function rectanglesInBoundingBox(boundingBox, rectangleWidth, rectangleHeight) {
	const rectangles = [];

	// Calculate the number of rectangles that can fit in each dimension
	const rowCount = Math.floor((boundingBox.maxY - boundingBox.minY) / rectangleHeight);
	const colCount = Math.floor((boundingBox.maxX - boundingBox.minX) / rectangleWidth);
	console.log(rowCount, colCount);
	// Generate the coordinates for each rectangle
	for (let row = 0; row < rowCount; row++) {
		for (let col = 0; col < colCount; col++) {
			const x = boundingBox.minX + col * rectangleWidth;
			const y = boundingBox.minY + row * rectangleHeight;
			// console.log(x, col, y, row);
			rectangles.push([
				[x, y],
				[x + rectangleWidth, y],
				[x + rectangleWidth, y + rectangleHeight],
				[x, y + rectangleHeight]
			]);
		}
	}

	return rectangles;
}

function isRectangleInsidePolygon(polygon, rectangle) {
	return rectangle.every((point) => {
		console.log(point, isPointInsidePolygon(polygon, point));
		return isPointInsidePolygon(polygon, point);
	});
}

function isPointInsidePolygon(polygon, point) {
	const x = point[0];
	const y = point[1];
	let isInside = false;

	let j = polygon.length - 1;
	for (let i = 0; i < polygon.length; i++) {
		const xi = polygon[i][0];
		const yi = polygon[i][1];
		const xj = polygon[j][0];
		const yj = polygon[j][1];

		const intersect = yi > y !== yj > y && x < ((xj - xi) * (y - yi)) / (yj - yi) + xi;
		if (intersect) {
			isInside = !isInside;
		}
		j = i;
	}

	return isInside;
}

const bb = calculateBoundingBox(points);
console.log(bb);
const rectangles = rectanglesInBoundingBox(bb, 20, 20);

const distances = points
	.map((point, index) => {
		if (index < points.length - 1) {
			return distance(point, points[index + 1]);
		}
	})
	.filter((point) => point !== undefined);
// console.log(distances, distances.length);

// get all rectuangles that are inside the polygon
const insideRectangles = rectangles.filter((rectangle) =>
	isRectangleInsidePolygon(points, rectangle)
);
// console.log(rectangles)
console.log(rectangles.length, insideRectangles.length);

function createSVG(coordinates, width, height) {
	const svgHeader = `<svg xmlns="http://www.w3.org/2000/svg" width="${width}" height="${height}">\n`;
	const svgFooter = '</svg>';

	const coordArray = coordinates.split(' ').map(parseFloat);

	// Find min and max values for x and y
	const minX = Math.min(...coordArray.filter((_, i) => i % 2 === 0));
	const maxX = Math.max(...coordArray.filter((_, i) => i % 2 === 0));
	const minY = Math.min(...coordArray.filter((_, i) => i % 2 === 1));
	const maxY = Math.max(...coordArray.filter((_, i) => i % 2 === 1));

	// Calculate scale factors for x and y
	const scaleX = (width - 10) / (maxX - minX);
	const scaleY = (height - 10) / (maxY - minY);

	// Transform the coordinates to fit within the SVG dimensions
	const transformedPoints = coordArray
		.map((coord, index) => {
			const scaledCoord = index % 2 === 0 ? (coord - minX) * scaleX : (coord - minY) * scaleY;
			return scaledCoord;
		})
		.map((coord, index) => (index % 2 === 0 ? `${coord},` : `${coord}`))
		.join('');

	const polygon = `<polygon points="${transformedPoints}" fill="blue" stroke="black" stroke-width="1"/>\n`;

	return svgHeader + polygon + svgFooter;
}

console.log(createSVG(stringPoints, 1000, 1000));
function polygonArea(polygon) {
	let area = 0;
	const n = polygon.length;

	for (let i = 0; i < n - 1; i++) {
		area += polygon[i][0] * polygon[i + 1][1] - polygon[i + 1][0] * polygon[i][1];
	}
	area += polygon[n - 1][0] * polygon[0][1] - polygon[0][0] * polygon[n - 1][1];

	return Math.abs(area) / 2;
}


const area = polygonArea(points);
console.log(area);