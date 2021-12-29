#[derive(Debug, Clone, PartialEq, Eq)]
pub struct LogicAction {
    pub target: String,
    pub source_left: String,
    pub source_right: String,
    pub operation: String,
}

pub fn parse(lines: Vec<String>) -> Vec<LogicAction> {
    let actions: Vec<LogicAction> = lines
        .iter()
        .map(|line| {
            let split = line
                .split(" -> ")
                .map(|x| x.to_string())
                .collect::<Vec<String>>();

            let mut parse = (String::new(), String::new(), String::new(), String::new());
            let source: Vec<String> = split[0].split(" ").map(|x| x.to_string()).collect();
            let target = split[1].clone();

            match source.len() {
                1 => {
                    parse = (
                        "ASSIGN".to_string(),
                        target.clone(),
                        source[0].clone(),
                        String::new(),
                    )
                }
                2 => parse = (source[0].clone(), target, String::new(), source[1].clone()),
                3 => {
                    parse = (
                        source[1].clone(),
                        target.clone(),
                        source[0].clone(),
                        source[2].clone(),
                    )
                }
                _ => {}
            }

            LogicAction {
                target: parse.1,
                operation: parse.0,
                source_left: parse.2,
                source_right: parse.3,
            }
        })
        .collect();

    actions
}
