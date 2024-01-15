package main

import "fmt"

// AddressBookNode represents a node in the address book BST.
type AddressBookNode struct {
 Name        string
 ContactInfo string
 Left        *AddressBookNode
 Right       *AddressBookNode
}

// InsertContact adds a new contact to the address book BST.
func (n *AddressBookNode) InsertContact(name, contactInfo string) *AddressBookNode {
 if n == nil {
  return &AddressBookNode{Name: name, ContactInfo: contactInfo, Left: nil, Right: nil}
 }

 if name < n.Name {
  n.Left = n.Left.InsertContact(name, contactInfo)
 } else if name > n.Name {
  n.Right = n.Right.InsertContact(name, contactInfo)
 }

 return n
}

// SearchContact searches for a contact in the address book BST.
func (n *AddressBookNode) SearchContact(name string) (string, bool) {
 if n == nil {
  return "", false
 }

 if name == n.Name {
  return n.ContactInfo, true
 }

 if name < n.Name {
  return n.Left.SearchContact(name)
 }
 return n.Right.SearchContact(name)
}

// DeleteContact removes a contact from the address book BST.
func (n *AddressBookNode) DeleteContact(name string) *AddressBookNode {
 if n == nil {
  return nil
 }

 if name < n.Name {
  n.Left = n.Left.DeleteContact(name)
 } else if name > n.Name {
  n.Right = n.Right.DeleteContact(name)
 } else {
  // Node with the contact to be removed found

  // Case 1: Node with only one child or no child
  if n.Left == nil {
   return n.Right
  } else if n.Right == nil {
   return n.Left
  }

  // Case 2: Node with two children
  // Find the smallest node in the right subtree (in-order successor)
  smallest := n.Right.FindMin()
  n.Name = smallest.Name
  n.ContactInfo = smallest.ContactInfo
  n.Right = n.Right.DeleteContact(smallest.Name)
 }

 return n
}

// Height returns the height of the address book BST.
func (n *AddressBookNode) Height() int {
 if n == nil {
  return 0
 }

 leftHeight := n.Left.Height()
 rightHeight := n.Right.Height()

 // Height of the tree is the maximum of left and right subtree heights, plus 1 for the current node
 if leftHeight > rightHeight {
  return leftHeight + 1
 }
 return rightHeight + 1
}

// findMin returns the node with the minimum value in the binary search tree.
func (n *AddressBookNode) FindMin() *AddressBookNode {
 current := n

 // Traverse the left subtree until the leftmost node is reached
 for current.Left != nil {
  current = current.Left
 }

 return current
}

func main() {
 // Create an empty address book BST
 var addressBook *AddressBookNode

 // Insert contacts
 addressBook = addressBook.InsertContact("John Doe", "123-456-7890")
 addressBook = addressBook.InsertContact("Alice Smith", "987-654-3210")
 addressBook = addressBook.InsertContact("Bob Johnson", "555-123-4567")

 // Search for a contact
 contactInfo, found := addressBook.SearchContact("Alice Smith")
 if found {
  fmt.Printf("Contact found: %s\n", contactInfo)
 } else {
  fmt.Println("Contact not found.")
 }

 // Delete a contact
 addressBook = addressBook.DeleteContact("Alice Smith")

 // Calculate the height of the address book
 treeHeight := addressBook.Height()
 fmt.Printf("Address book height: %d\n", treeHeight)
}
