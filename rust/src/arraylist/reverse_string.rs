pub fn reverse_string(s: &mut Vec<char>) {
    let mut p1 = 0_usize;
    let mut p2 = s.len()-1;
    while p1 < p2 {
        let tmp = s[p1];
        s[p1] = s[p2];
        s[p2] = tmp;
        p1 += 1;
        p2 -= 1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_reverse_string() {
        let mut s = vec!['h', 'e', 'l', 'l', 'o'];
        reverse_string(&mut s);
        assert_eq!(vec!['o', 'l', 'l', 'e', 'h'], s)
    }
}