package templates

const MintKibblesTemplate = `
import FungibleToken from 0x{{.FungibleTokenAddress}} 
import Kibble from 0x{{.KibbleTokenAddress}}

transaction(recipient: Address, amount: UFix64) {
    let tokenAdmin: &Kibble.Administrator
    let tokenReceiver: &{FungibleToken.Receiver}

    prepare(signer: AuthAccount) {
        self.tokenAdmin = signer
        .borrow<&Kibble.Administrator>(from: /storage/KibbleAdmin) 
        ?? panic("Signer is not the token admin")

        self.tokenReceiver = getAccount(recipient)
        .getCapability(/public/KibbleReceiver)!
        .borrow<&{FungibleToken.Receiver}>()
        ?? panic("Unable to borrow receiver reference")
    }

    execute {
        let minter <- self.tokenAdmin.createNewMinter(allowedAmount: amount)
        let mintedVault <- minter.mintTokens(amount: amount)

        self.tokenReceiver.deposit(from: <-mintedVault)

        destroy minter
    }
}
`
