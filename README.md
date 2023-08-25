# Averveil

![hackmd-github-sync-badge](https://hackmd.io/6kXvWxo7RXKp1RcMFpWlbQ/badge)


## Overview
Within the realm of Hyperversal Blocks, Averveil is a multiverse comprising of multiple smaller universes(building blocks), each having their own goals and ambitions and could be plugged to any other system.

### Miniverses
Averveil primarily consists of these miniverses:
- Escrow
- ZkInsights
- ZkLabels
- ZkTracker
- PillarBox

### Wormholes
Wormholes are the systems that will allow multiple universes or multiverses to connect or communicate with each other:
- ChainRegister
- DAO

## Usecases of miniverses and wormholes in Averveil

### Escrow
#### Overview
A blockchain based temporary legal arrangement between 2 transacting parties where a third party holds the financial payment.This innovative solution aims to provide a secure and transparent platform for facilitating transactions. By harnessing the power of blockchain technology, the app ensures that transactions are executed only when predetermined conditions are met, eliminating the need for traditional intermediaries and fostering a high level of trust among participants. This project aligns with the ethos of decentralization and self-executing smart contracts, contributing to a more efficient and equitable digital economy.

#### Operational Framework and Technical Mechanics
The project entails the creation of a sophisticated smart contract designed to seamlessly manage transactions involving parties. These transactions include both outbound payments and inbound receipts. An event trigger, such as the "buy now" action, prompts the deduction of the required payment from Party1's wallet. This payment is then securely routed to a designated Escrow account.

The Escrow account functions as a vigilant intermediary, patiently awaiting the satisfaction of predetermined conditions. Once these conditions are met, the Escrow account dutifully facilitates the transfer of funds to Party2, culminating in the successful execution of the transaction.

A notable aspect is the Escrow account's role in collecting a commission for its services. This commission is thoughtfully directed to a separate wallet, enhancing the financial dynamics of the system.

Upon the accomplishment of the Minimum Viable Product (MVP) milestone for "Averveil," the commission takes on a new purpose. It becomes a source of equitable allocation under the purview of a Decentralized Autonomous Organization (DAO) managed by Chain Registrars (Custodians). These DAO custodians oversee the judicious distribution of funds, ensuring a fair and transparent disbursement process.

### ZkInsights
ZkInsights aggregates information from diverse sources. Data is subjected to encryption protocols to formulate an encrypted dataset. The application then employs a zero-knowledge proof methodology to facilitate the dissemination of this encrypted information. The overarching goal is to enable secure and confidential data sharing among interconnected applications. By leveraging advanced zero-knowledge techniques, the application ensures that data can be shared without revealing its underlying contents, thereby safeguarding privacy and confidentiality.

#### Operational Framework and Technical Mechanics
The functioning process of ZkInsights involves the utilization of input structured in JSON format, presented as an object. This input is systematically processed within the framework of zero-knowledge proofs (ZKPs), ensuring that delicate details remain confidential. For instance, the ZKP technique can be utilized to validate an assertion, such as confirming an individual's legal age without revealing the actual age, accomplished through zk-SNARKs.

The incoming data, in its covert form, is skillfully parsed for further processing. The object's keys consistently adhere to plain English, facilitating a coherent and easily understandable structure.

This compilation of information then serves as the cornerstone for constructing a comprehensive dataset. This dataset is subsequently poised for distribution to marketing or analytics entities, functioning as a valuable resource that comes at a designated cost.

A fundamental aspect of the user experience is the empowerment extended to users regarding data sharing. Individuals possess the prerogative to withhold their data, exercising their right to decline participation. This pivotal choice is fortified through a verification mechanism, which relies on user signatures. These signatures are meticulously validated through dedicated wallet transactions, reinforcing the integrity of the user's decision.

### ZKLabels
ZkLabels streamlines handling confidential data with zero-knowledge proofs, converting it into label-like structures for applications like shipping management. It enables planned sharing, notifying designated services promptly. Information gathered by ZKLabels can be destroyed based on predefined conditions or can be used to generate insights with explicit permission.

#### Operational Framework and Technical Mechanics
ZkLabels comes into play by handling the information that has already been privacy-protected using zero-knowledge proofs (ZKPs). This information is transformed into specific pre-defined data structures, resembling labels, which can serve various purposes like organizing shipping details.

ZkLabel's role extends to guaranteeing that the sharing of information can be planned in advance. This means that if a particular service requires early notification about the presence of a specific label, it can effortlessly keep track using ZkInsights.

In essence, ZkLabels works behind the scenes to arrange secure information, turning it into handy labels, and making sure that the process of sharing data is streamlined and well-managed.

In future, the goal is to integrate AI so that instead of a pre-defined data structure, it can take in any information and convert it into any necessary structure.

### ZKTracker
ZkTracker is a decentralized service that links encrypted and zk-proven addresses/locations (derived from GPS sensors) with registered shipping services. This connection occurs via blockchain. Couriers can scan ZkLabels-generated labels, revealing pinpoint GPS locations. Information gradually unfolds based on conditions; e.g., riders get broader details until meeting location requirements. This controlled disclosure maintains privacy, fostering efficient and secure distribution.

Information gathered by ZKTracker can be destroyed based on predefined conditions or can be used to generate insights with explicit permission from respective parties.

#### Operational Framework and Technical Mechanics
ZkTracker plays a pivotal role as a decentralized nexus, effectively connecting encrypted and zk-proven addresses/locations sourced from GPS sensors with an extensive network of registered shipping services. All interactions are underpinned by the secure framework of blockchain technology.

Consider a scenario where a courier rider, representing a courier company, interacts with the system. By scanning a label generated by ZkLabels or a compatible service, the rider gains access to precise GPS coordinates. This seamless process significantly enhances the ability to trace and monitor shipments.

The distinctive feature of ZkTracker lies in its nuanced information dissemination. The data flow unfolds progressively, offering insights into the shipment's trajectory. For instance, when a tracker initiates its journey from a different country, only generalized country information is accessible. As it reaches specific locales such as cities or towns, the granularity of information gradually increases, culminating in the provision of highly precise location details.

Notably, the privacy mechanisms employed work in dual directions. While local riders can access evolving data about their area, they remain unable to backtrack and identify the parcel's originating location. This two-way privacy measure ensures a balanced and secure system where sensitive details are safeguarded. This approach reinforces secure and efficient distribution while maintaining a robust layer of privacy protection.

ZkTracker can also provide analytics to other services like ZkInsights.

### PillarBox
Pillarbox, an essential network of sensors, is intricately woven into the fabric of Averveil. These sensors are indispensable to the Averveil system, actively engaged in each step of the order process and providing encrypted data crucial for various services. Averveil's functionality relies on this dynamic interaction, and in turn, Pillarbox gains significance from Averveil's integration. From validating labels to chronicling shipping progression and acting as a reservoir of contextual data from ZkTracker and ZkLabels, Pillarbox's symbiotic relationship with Averveil is the cornerstone of operational effectiveness and advancement in the blockchain domain. This system primarily gathers focused user-centric and product-centric information instead of general insights provided by ZkInsights.

#### Operational Framework and Technical Mechanics
Pillarbox, a network of intricately designed sensors, constitutes a foundational component of the Averveil system. These sensors are seamlessly registered within the Averveil architecture, playing a pivotal role in maintaining the system's functionality.

Every step of the order placement process involves direct communication with these sensors. They are an indispensable part of the process, ensuring that Averveil remains fully operational. The information gathered by these sensors is transmitted in a secure and encrypted format, serving as a valuable resource for various services within the Averveil ecosystem.

Importantly, the interdependence is mutual. Pillarbox sensors are indispensable to Averveil, while the system's utilization enhances the capabilities of Pillarbox. During label generation, these sensors not only provide the vital signature for label validation but also chronicle the evolving state of shipping.

Similarly, during ZkTracker's information dissemination, Pillarbox serves as a reservoir of comprehensive insights. This repository extends beyond what is shared through ZkTracker, capturing a wider spectrum of shipping-related data. This broader scope is attributed to Pillarbox's foundational role in initiating the process by providing the original address.

In essence, Pillarbox's connection is essential, serving as a linchpin in the Averveil ecosystem. Its continuous interaction with Averveil is fundamental; the system relies on this integration for its core functionality. 

### ChainRegister
#### Overview
Chain Register is a decentralized node-based system designed to facilitate the seamless operation of various services within the Averveil multiverse as well as other future multiverses. Node operators play a central role by booting up nodes that host a range of services, including those mentioned earlier (Escrow, ZkInsights, ZkLabels, ZkTracker, PillarBox) as well as additional essential services like Swarm and IPFS. These nodes collectively form the backbone of the Averveil ecosystem.

#### Node Setup and Service Deployment
Node operators take responsibility for initializing nodes within the Chain Register system. Upon booting up, nodes automatically instantiate instances of the specified services, ensuring their availability to network participants. This automated process includes the setup of in-memory databases and other requisite services. Node operators' participation in the system is incentivized through a rewards mechanism.

#### Rewards Mechanism
The rewards bestowed upon node operators are determined through a dual mechanism incorporating both traditional staking and liquidity pools. Node operators stake a certain amount of tokens to participate in the system. This stake signifies their commitment to the network's stability and security. Additionally, liquidity pools play a pivotal role by providing necessary liquidity to the ecosystem, ensuring its self-sufficiency. Rewards are proportionally allocated based on the staking amount and the contribution to liquidity pools.

#### DAO-Managed Liquidity Pools
To uphold transparency and trustlessness, liquidity pools are governed by a Decentralized Autonomous Organization (DAO). The DAO oversees the operation of these pools, ensuring that rewards are transferred fairly and efficiently. This collective management enhances the integrity of the liquidity provisioning process and safeguards against centralization. DAOs will also manage the funding of other services.

#### Pluggable Architecture
The Chain Register system is designed with a pluggable architecture, enabling easy integration of additional services. As the Averveil multiverse evolves and new services emerge, node operators can seamlessly incorporate these services into the existing infrastructure. This flexibility ensures that the system remains adaptable and future-proof.

#### Multiverse Preference and Reward Customization
Node operators are empowered to choose the specific multiverses they wish to connect with. This multiverse preference impacts their potential rewards. Operators that contribute to and serve the preferred multiverses receive rewards accordingly. This personalized approach fosters engagement and aligns incentives with operators' strategic choices.

#### Data Collection, Storage, and Retrieval
Node operators play a pivotal role in collecting, storing, and retrieving data from various services and smart contracts within the Averveil ecosystem. Data is securely stored using Swarm, a decentralized storage protocol. Operators ensure the availability and integrity of information, facilitating efficient retrieval and sharing as needed.

The Chain Register system acts as a cohesive framework that harmonizes the operation of services, incentivizes node operators, and drives the sustainability and expansion of the Averveil multiverse.

## Data that be stored on Swarm:
Here are the points where data storage is involved where Swarm will be used as storage medium:

1. **Escrow:**
   - Data Collection: Transaction data, cost of transaction, sensor information, participant identities.
   - Commission Record: Collection of commission in a separate wallet.
   - Financial information and blockchain state.

2. **ZkInsights:**
   - Data Aggregation: Gathering data from diverse sources.
   - Encrypted Dataset: Formulation of an encrypted dataset.
   - Data Distribution: Distribution of the dataset to marketing or analytics entities and/or collecting their data.

3. **ZkLabels:**
   - Privacy-Protected Data Handling: Transformation of already privacy-protected data into predefined structures.
   - Planned Information Sharing: Management of planned sharing of information.
   - AI Integration (Future): Goal to integrate AI for flexible structure conversion.

4. **ZkTracker:**
   - Encrypted Address Integration: Linking encrypted and zk-proven addresses/locations with shipping services.
   - Progressive Information Dissemination: Controlled disclosure of information as shipment progresses.
   - Analytics Provision: Potential provision of analytics to other services like ZkInsights.

5. **PillarBox:**
   - Sensor Data Collection: Gathering of data such as location, timestamps, and status throughout the shipping process.
   - Label Generation Support: Provision of signatures for validating labels generated by ZkLabels.
   - Shipping State Chronicle: Recording and chronicling the evolving state of shipping.

## Quaterly Milestones Breakdown
### MVP
#### Milestones for MVP
##### Milestone Q1: Project Initiation, Planning and Infrastructure Setup
- [ ] Identify key features and functionalities of each service.
- [ ] Break down high-level requirements for each services in scope of mvp for Sprint Planning.
- [ ] Identify tasks within each sprint.
- [ ] Set up communication channels for collaboration.
- [ ] Prepare project documentation.
- [ ] Develop a secure node boot-up mechanism for Chain Register.
- [ ] Create scripts for automated deployment of node instances.
- [ ] Implement an MVP of each service that will be extended in Node. 
- [ ] Set up and configure Swarm and IPFS for data storage.

##### Milestone Q2 and Q3: Contract and Services
- [ ] Design the ERC20 Contract.
- [ ] Define tokenomics and DAO structure.
- [ ] Design the structure of liquidity pools.
- [ ] Identify and implement mock contracts for all services.
- [ ] Implement smart contracts for liquidity pool management.
- [ ] Implement the Mock Services idenfitifed in Sprints with value objects and data transfer objects.
- [ ] Re-configure the node structure based on the newly identified mocks of other services with appropriate data transfer objects and value objects.

##### Milestone Q4: TBD