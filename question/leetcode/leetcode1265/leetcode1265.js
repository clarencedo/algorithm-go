function printLinkedListInReverse(head: ImmutableListNode) {
  if (head != null) {
    printLinkedListInReverse(head.getNext());
    head.printValue();
  }
};