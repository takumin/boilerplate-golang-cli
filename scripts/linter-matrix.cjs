module.exports = ({ json }) => {
	const needs = JSON.parse(json);
	var targets = [];
	if (needs["github-actions"] === "true") {
		targets.push("actionlint");
	}
	return targets;
};
