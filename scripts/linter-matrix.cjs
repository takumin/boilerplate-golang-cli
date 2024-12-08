module.exports = ({ needs }) => {
	var targets = [];
	if (needs["github-actions"] === "true") {
		targets.push("actionlint");
	}
	return targets;
};
