use std::collections::{HashMap, VecDeque};

use crate::day7::input::LogicAction;
use helpers::{ChallengeSolution, Solution};

mod input;

/// Parse value will return the value if it has already been defined in the map
/// otherwise it will return no value, (None). In which cause the operation should
/// be ignored.
///
/// # Arguments
///
/// * `value`:
/// * `logic_map`:
///
/// returns: Option<u16>
fn parse_value(value: &String, logic_map: &HashMap<String, u16>) -> Option<u16> {
    match value.parse::<u16>() {
        Ok(value) => Some(value),
        Err(_) => {
            if let Some(logic_value) = logic_map.get(value) {
                return Some(*logic_value);
            }
            None
        }
    }
}

fn not(action: &LogicAction, logic_map: &mut HashMap<String, u16>) -> bool {
    let value = parse_value(&action.source_right, logic_map);

    if let Some(parse_value) = value {
        logic_map.insert(action.target.clone(), !parse_value);
        return true;
    }

    false
}

fn assign(action: &LogicAction, logic_map: &mut HashMap<String, u16>) -> bool {
    let value = parse_value(&action.source_left, logic_map);

    if let Some(parse_value) = value {
        logic_map.insert(action.target.clone(), parse_value);
        return true;
    }
    false
}

fn and_or(
    left: Option<u16>,
    right: Option<u16>,
    target: &String,
    logic_map: &mut HashMap<String, u16>,
    or: bool,
) -> bool {
    if let Some(parsed_left) = left {
        if let Some(parsed_right) = right {
            let mut value = parsed_left & parsed_right;

            if or {
                value = parsed_left | parsed_right;
            }

            logic_map.insert(target.clone(), value);
            return true;
        }
    }

    false
}

fn shift(
    left: Option<u16>,
    right: Option<u16>,
    target: &String,
    logic_map: &mut HashMap<String, u16>,
    right_shift: bool,
) -> bool {
    if let Some(parsed_left) = left {
        if let Some(parsed_right) = right {
            return if right_shift {
                logic_map.insert(target.clone(), parsed_left >> parsed_right);
                true
            } else {
                logic_map.insert(target.clone(), parsed_left << parsed_right);
                true
            };
        }
    }

    false
}

fn solve(logic_actions: &Vec<LogicAction>) -> String {
    // the logical hash map used to keep track of the current
    // value of the bit operation on that value.
    let mut logic_actions_map: HashMap<String, u16> = HashMap::new();
    let mut queue = VecDeque::from(logic_actions.clone());

    while queue.len() != 0 {
        let action = queue.pop_front().unwrap();

        let left = parse_value(&action.source_left, &logic_actions_map);
        let right = parse_value(&action.source_right, &logic_actions_map);

        let ok = match action.operation.as_str() {
            "ASSIGN" => assign(&action, &mut logic_actions_map),
            "NOT" => not(&action, &mut logic_actions_map),
            "AND" => and_or(left, right, &action.target, &mut logic_actions_map, false),
            "OR" => and_or(left, right, &action.target, &mut logic_actions_map, true),
            "RSHIFT" => shift(left, right, &action.target, &mut logic_actions_map, true),
            "LSHIFT" => shift(left, right, &action.target, &mut logic_actions_map, false),
            _ => false,
        };

        if !ok {
            queue.push_back(action)
        }
    }

    logic_actions_map.get("a").unwrap().to_string()
}

pub(crate) struct Challenge {}

impl ChallengeSolution for Challenge {
    fn solve(&self, input: Vec<String>) -> Solution {
        let mut logic_actions = input::parse(input);
        let part_one = solve(&logic_actions);

        for x in logic_actions.iter_mut() {
            if x.operation == "ASSIGN" && x.target == "b" {
                x.source_left = part_one.to_string();
            }
        }

        let part_two = solve(&logic_actions);

        Solution {
            part1: part_one,
            part2: part_two,
        }
    }

    fn year(&self) -> i32 {
        2015
    }

    fn day(&self) -> i32 {
        7
    }

    fn example(&self) -> bool {
        false
    }
}
