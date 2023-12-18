pub fn longest_palindrome(s: String) -> String {
    let mut res = "";
    for i in 1..s.len() {
        let s1 = palindrome(&s, i as isize , i as isize);
        let s2 = palindrome(&s, i as isize, (i+1) as isize);

        if s1.len() > res.len() {
            res = s1;
        }
        if s2.len() > res.len() {
            res = s2;
        }
    }

    res.to_string()
}

fn palindrome(s: &str, l: isize, r: isize) -> &str {
    let mut l = l;
    let mut r = r;
    let bytes = s.as_bytes();
    while l >= 0 && r < s.len() as isize && bytes[l as usize] == bytes[r as usize] {
        l -= 1;
        r += 1;
    }

    &s[(l+1) as usize..r as usize]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest_palindrome() {
        let s = "babad";
        let res = longest_palindrome(s.to_string());
        assert_eq!("bab".to_string(), res);

        let s = "cbbd";
        let res = longest_palindrome(s.to_string());
        assert_eq!("bb".to_string(), res);
    }
}
