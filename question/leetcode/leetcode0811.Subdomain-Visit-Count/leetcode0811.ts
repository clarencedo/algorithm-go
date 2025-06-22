function subdomainVisits(cpdomains: string[]): string[] {
	const ans: string[] = [];
	let countMap = new Map<string, number>();

	for (const cpdomain of cpdomains) {
		const [freqStr, fulldomain] = cpdomain.split(" ");
		const freq = Number(freqStr);
		const domains = fulldomain.split(".");

		for (let i = 0; i < domains.length; i++) {
			const subdomain = domains.slice(i).join(".");
			countMap.set(subdomain, (countMap.get(subdomain) || 0) + freq);
		}
	}

	for (const [domain, freq] of countMap.entries()) {
		ans.push(`${freq} ${domain}`);
	}

	return ans;
};