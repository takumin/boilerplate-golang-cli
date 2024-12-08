module.exports = (changes) => {
	var targets = [];
	if (changes["github-actions"] === "true") {
		targets.push("actionlint");
	}
	return targets;
};
