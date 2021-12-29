#[cfg(test)]
mod test_decode {
    use helpers::ChallengeSolution;

    use crate::day8::{parse_string_length, Challenge};

    #[test]
    fn parse_string_length_test() {
        let letter_input = vec![
            ("\"\"", 2, 0),
            ("\"abc\"", 5, 3),
            ("\"aaa\\\"aaa\"", 10, 7),
            ("\"\\x27\"", 6, 1),
        ];

        for x in &letter_input {
            let input = x.0.to_string();
            let resp = parse_string_length(input);

            assert_eq!(resp.0, x.1);
            assert_eq!(resp.1, x.2);
        }

        let solve_input = letter_input
            .into_iter()
            .map(|x| x.0.to_string())
            .collect::<Vec<String>>();

        let challenge = Challenge {};
        assert_eq!(challenge.solve(solve_input).part1, 12.to_string());
    }
}
