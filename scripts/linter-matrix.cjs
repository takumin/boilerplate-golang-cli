module.exports = (needs) => {
	console.log(needs);

	var targets = [];
	if (needs["github-actions"] === "true") {
		targets.push("actionlint");
	}
	return targets;
};
