// TODO: 理解还不够透彻
class TrieNode {
	children: Map<string, TrieNode>; //每个字符到对应子节点
	isEnd: boolean;// 标记当前节点是否是某个单词的结尾

	constructor() {
		this.children = new Map();
		this.isEnd = false;
	}
}

class Trie {
	private root: TrieNode;

	constructor() {
		this.root = new TrieNode();
	}

	insert(word: string): void {
		let node = this.root;
		for (const ch of word) {
			if (!node.children.has(ch)) {
				node.children.set(ch, new TrieNode());
			}
			node = node.children.get(ch)!;
		}
		node.isEnd = true;
	}

	search(word: string): boolean {
		const node = this.searchPrefix(word);
		return node !== null && node.isEnd;
	}

	startsWith(prefix: string): boolean {
		return this.searchPrefix(prefix) !== null;
	}

	private searchPrefix(prefix: string): TrieNode | null {
		let node = this.root;
		for (const ch of prefix) {
			if (!node.children.has(ch)) {
				return null;
			}
			node = node.children.get(ch)!;
		}
		return node;
	}
}
