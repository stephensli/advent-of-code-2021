use helpers::ChallengeSolution;
use solutions::*;

mod solutions;

fn main() {
    let days: Vec<Box<dyn ChallengeSolution>> = vec![
        Box::new(day7::Challenge {}),
        Box::new(day8::Challenge {}),
    ];

    // latest
    let challenge_number = days.len() - 1;

    // selective
    // let challenge_number = 9 - 7;

    solve(&days[challenge_number]);
}

pub fn solve(challenge: &Box<dyn ChallengeSolution>) {
    let mut setup = helpers::setup(challenge.year(), challenge.day(), challenge.example());
    let lines = helpers::get_lines(&setup.input_file_path);
    setup.solution = challenge.solve(lines);
}
