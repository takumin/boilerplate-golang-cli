module.exports = (changes) => {
	var targets = ["editorconfig-checker"];
	if (actionlint(changes)) {
		targets.push("actionlint");
	}
	if (json(changes)) {
		targets.push("gjson-validate");
	}
	if (yaml(changes)) {
		targets.push("gyaml-validate");
	}
	if (shell(changes)) {
		targets.push("shellcheck");
		targets.push("shfmt");
	}
	if (golang(changes)) {
		targets.push("gofmt");
		targets.push("gosec");
		targets.push("govet");
		targets.push("staticcheck");
	}
	return targets;
};

function actionlint(changes) {
	if (
		changes["aqua"] === "true" ||
		changes["reviewdog"] === "true" ||
		changes["github-actions"] === "true"
	) {
		return true;
	} else {
		return false;
	}
}

function yaml(changes) {
	if (
		changes["aqua"] === "true" ||
		changes["reviewdog"] === "true" ||
		changes["yaml"] === "true"
	) {
		return true;
	} else {
		return false;
	}
}

function json(changes) {
	if (
		changes["aqua"] === "true" ||
		changes["reviewdog"] === "true" ||
		changes["json"] === "true"
	) {
		return true;
	} else {
		return false;
	}
}

function shell(changes) {
	if (
		changes["aqua"] === "true" ||
		changes["reviewdog"] === "true" ||
		changes["shell"] === "true"
	) {
		return true;
	} else {
		return false;
	}
}

function golang(changes) {
	if (
		changes["aqua"] === "true" ||
		changes["reviewdog"] === "true" ||
		changes["go"] === "true"
	) {
		return true;
	} else {
		return false;
	}
}
