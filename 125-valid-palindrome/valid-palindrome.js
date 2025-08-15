/**
 * @param {string} s
 * @return {boolean}
 */
 function isAlphaNum (c) {
    return (
        (c >= 'a' && c <= 'z') ||
        (c >= 'A' && c <= 'Z') ||
        (c >= '0' && c <= '9')
    )
 }
var isPalindrome = function(s) {
    let fr = 0;
    let bc = s.length - 1;

    while (fr < bc) {
        while (fr < bc && !isAlphaNum(s[fr])) {
            fr++
        }
        while (fr < bc && !isAlphaNum(s[bc])) {
            bc--
        }

        if (s[fr].toLowerCase() !== s[bc].toLowerCase()) {
            return false
        }
        fr++;
        bc--
    }    

    return true
};