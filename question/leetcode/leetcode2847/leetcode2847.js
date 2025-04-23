/**
 * @param {number} n
 * @return {string}
 */
// js number会在大数的时候丢失精度
var smallestNumber = function (n) {
	if (n < 10) {
		return n.toString();
	}

	let factors = [];
	for (let i = 9; i >= 2; i--) {
		while (n % i === 0) {
			factors.push(i);
			n = n / i;
		}
	}

	if (n > 1) {
		return "-1";
	}

	return factors.sort((a, b) => a - b).join('');
};
