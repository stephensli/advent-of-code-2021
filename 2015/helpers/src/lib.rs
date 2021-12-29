use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::{Path, PathBuf};
use std::time::Instant;

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Solution {
    pub part1: String,
    pub part2: String,
}

pub trait ChallengeSolution {
    fn solve(&self, input: Vec<String>) -> Solution;
    fn year(&self) -> i32;
    fn day(&self) -> i32;
    fn example(&self) -> bool;
}

/// Helper to read lines into a iterator
///
/// https://doc.rust-lang.org/stable/rust-by-example/std_misc/file/read_lines.html
///
/// # Arguments
///
/// * `filename`: The source file path.
///
/// returns: Result<Lines<BufReader<File>>, Error>
///
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

pub fn get_lines(path: &Path) -> Vec<String> {
    read_lines(path)
        .unwrap()
        .map(|x| x.unwrap())
        .collect::<Vec<String>>()
}

fn print_example_flag() {
    println!("\n########### ###########\n# EXAMPLE # # EXAMPLE #\n########### ###########\n");
}

pub struct Setup {
    year: i32,
    day: i32,
    now: Instant,
    example: bool,
    pub solution: Solution,
    pub input_file_path: PathBuf,
}

impl Drop for Setup {
    /// drop will result in the AOC print out of the solutions
    /// and execution time of the run. A example flag will be
    /// shown if required.
    fn drop(&mut self) {
        if self.solution.part1.trim() == "" {
            self.solution.part1 = "Not Provided".to_string();
        }

        if self.solution.part2.trim() == "" {
            self.solution.part2 = "Not Provided".to_string();
        }

        if self.example {
            print_example_flag()
        }

        println!(
            "\nyear:\t\t- {}\n\
            day:\t\t- {:02}\n\
            duration:\t- {:2?}\n\
            Part-One:\t- {}\n\
            Part-Two:\t- {}\n",
            self.year,
            self.day,
            self.now.elapsed(),
            self.solution.part1,
            self.solution.part2,
        );

        if self.example {
            print_example_flag()
        }
    }
}

/// Setup configures the day for running and allowing the solution to be
/// printed as expected. This is a standard way of outputting content.
///
/// It will also change the working directory to the source folder.
///
/// # Arguments
///
/// * `year`: The year in which the aoc was created.
/// * `day`: The day that is being exucted.
/// * `example`: If its the example that should be running.
///
/// returns: (String, Box<dyn Fn(Solution), Global>)
///
/// # Examples
/// ```
/// use helpers::Solution;
/// let completion = setup(2015, 1);
///
/// // code to complete day
///
/// completion(Solution{ part1: "part1".to_string(), part2: "part2".to_string()})
/// ```
pub fn setup(year: i32, day: i32, example: bool) -> Setup {
    let now = Instant::now();

    let current_directory = env::current_dir().unwrap();

    let mut target_path = format!("inputs/day{}", day);

    if example {
        target_path += ".example.txt"
    } else {
        target_path += ".txt"
    }

    let target_directory = Path::new(&target_path);
    let input_source_path = current_directory.join(target_directory);

    Setup {
        year,
        day,
        now,
        example,

        input_file_path: input_source_path,
        solution: Solution {
            part1: "".to_string(),
            part2: "".to_string(),
        },
    }
}
