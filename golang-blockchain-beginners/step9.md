# Create createNewBlock function

We need a function to handle the creation of a new block. This will take in the previous Block, a Voter, and a Candidate, and use them to create the new Block on the Blockchain.

Click to add this code to our editor.

<pre class="file" data-filename="main.go" data-target="append">
func createNewBlock(prevBlock Block, Voter, Candidate string) Block {
	t := time.Now()
	newBlock := Block{
		Timestamp: t.String(),
		Voter:     Voter,
		Candidate: Candidate,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = createHash(newBlock)
	return newBlock
}
</pre>

Let's continue to step 10.