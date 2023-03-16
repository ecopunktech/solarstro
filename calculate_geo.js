// 697985.83 4824277.34 697991.33 4824267.19 697987.76 4824265.25 697986.92 4824266.79 697981.56 4824263.88 697979.84 4824267.04 697980.72 4824267.52 697977.77 4824272.97 697985.83 4824277.34

const stringPoints =
	'697985.83 4824277.34 697991.33 4824267.19 697987.76 4824265.25 697986.92 4824266.79 697981.56 4824263.88 697979.84 4824267.04 697980.72 4824267.52 697977.77 4824272.97 697985.83 4824277.34';

const points = stringPoints
	.split(' ')
	.map((point, index) => {
		if (index % 2 === 0) {
			return [point, stringPoints.split(' ')[index + 1]];
		}
	})
	.filter((point) => point !== undefined);

const distance = (point1, point2) => {
	const x = point2[0] - point1[0];
	const y = point2[1] - point1[1];
	return Math.sqrt(x * x + y * y);
};

const distances = points.map((point, index) => {
    if (index < points.length - 1) {
        return distance(point, points[index + 1]);
    }
}).filter((point) => point !== undefined);
console.log(distances, distances.length);