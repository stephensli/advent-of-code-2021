use helpers::{ChallengeSolution, Solution};

#[cfg(test)]
mod test;

/// Returns the total character count and the total number of characters in
/// memory.
///
/// # Arguments
///
/// * `input`: The input string to determine the characters length.
///
/// returns: (usize, usize)
pub fn parse_string_length(input: String) -> (usize, usize) {
    let unquoted = enquote::unquote(&*input).unwrap();

    let total_characters = input.len();
    let total_memory_characters: usize = unquoted.chars().map(|x| x.len_utf16()).sum();

    (total_characters, total_memory_characters)
}

pub(crate) struct Challenge {}

impl ChallengeSolution for Challenge {
    fn solve(&self, input: Vec<String>) -> Solution {
        let mut count = 0;
        let mut count_two = 0;

        for x in input {
            let res = parse_string_length(x.clone());
            count += res.0 - res.1;

            // you should now encode each code representation as a new string and
            // find the number of characters of the new encoded representation,
            // including the surrounding double quotes.
            //
            // Just count the number of new characters being added which will just be
            // what is already escaped + the newly added outside escaping.
            //
            // since the total is minus the org string anyway this will be our value.
            count_two += 2 + x
                .clone()
                .chars()
                .filter(|x| *x == '\\' || *x == '"')
                .count();
        }

        Solution {
            part1: count.to_string(),
            part2: count_two.to_string(),
        }
    }

    fn year(&self) -> i32 {
        2015
    }

    fn day(&self) -> i32 {
        8
    }

    fn example(&self) -> bool {
        false
    }
}
