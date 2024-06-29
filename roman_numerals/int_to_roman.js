var intToRoman = function(num) {
    const numerals = new Map(
        [
            [1,"I"], [5,"V"],
            [10, "X"], [50, "L"],
            [100, "C"], [500, "D"],
            [1000, "M"]
        ]
    );
    let result = [];
    let position = 1;
    while (num > 0) {
        let digit = num % 10;
        if (digit < 4) {
            result.push(numerals.get(position).repeat(digit));
        }
        else if (digit === 4 || digit === 9) {
            result.push(numerals.get(position) + numerals.get((digit+1)*position));
        }
        else {
            result.push(numerals.get(5*position) + numerals.get(position).repeat(digit-5));
        }
        num = Math.floor(num / 10)
        position = position * 10;
    }
    return result.reverse().join('');
};


console.log(intToRoman(1), "I");
console.log(intToRoman(4), "IV");
console.log(intToRoman(46), "XLVI");
console.log(intToRoman(1994), "MCMXCIV");
// console.log(intToRoman(1), "I")