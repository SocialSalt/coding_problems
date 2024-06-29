var romanToInt = function(s) {
    const numerals = new Map(
        [
            ['I', 1], ['V', 5],
            ['X', 10], ['L', 50],
            ['C', 100], ['D', 500],
            ['M', 1000]
        ]
    );
    
    let last = numerals.get(s[0]);
    let result = 0;
    for (let i = 1; i < s.length; i++) {
        let curr = numerals.get(s[i]);
          if (curr > last && last > 0) {
              result += curr - last;
              last = 0;
          }
          else {
              result += last;
              last = curr;
          }
    }
    result += last;

    return result;
};

console.log(romanToInt("IV"), 4);
console.log(romanToInt("XIV"), 14);
console.log(romanToInt("XII"), 12);
console.log(romanToInt("XLVI"), 46);
console.log(romanToInt("VII"), 7);
console.log(romanToInt("III"), 3);
console.log(romanToInt("MCMX"), 1910);
console.log(romanToInt("MCMXCIV"), 1994);