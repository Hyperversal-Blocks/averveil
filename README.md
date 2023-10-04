# Averveil

![hackmd-github-sync-badge](https://hackmd.io/6kXvWxo7RXKp1RcMFpWlbQ/badge)

## Introduction
### Purpose
The purpose of this Software Requirement Specification (SRS) document is to provide a comprehensive and detailed description of the technical specifications and functionalities of "Averveil". This document aims to facilitate a clear understanding of the system's architecture, features, requirements, and the problems it intends to solve in the domain of data privacy and control. It serves as a foundational agreement between stakeholders and developers, ensuring that both parties have aligned expectations regarding the system’s capabilities, features, and limitations.

Averveil is envisioned as a revolutionary platform that addresses pressing concerns about data ownership, privacy, and security in our increasingly digital society. By leveraging cutting-edge technologies like blockchain, zero-knowledge proofs (ZKP), and verifiable credentials, Averveil aims to empower users with greater control over their data, facilitate secure and transparent transactions through an escrow system, and incentivize data sharing while preserving privacy.

This document also outlines the proposed strategies for engagement, growth, monetization, and marketing to ensure the success and sustainability of the Averveil platform.

### Scope
The scope of this SRS encompasses the development and deployment of the Averveil platform, aimed at addressing challenges in data privacy and decentralized transactions. The platform will include a variety of features and systems, focusing on user profiles, data sharing incentives, verifiable proofs, and secure blockchain transactions. Additionally, the project will involve the creation of user interfaces, integration with various software and hardware, and attention to non-functional requirements such as performance, usability, and security. The SRS also outlines strategies for engagement, growth, monetization, and marketing, along with plans for continuous improvement and adaptations to meet future needs and developments in the field.

### Glossary
ZKP - Zero knowledge proof
AV - Averveil
Node - One node amongst multiple nodes that can communicate with each other

### References
- [Authentication, Authorization, and Selective
  Disclosure for IoT data sharing using Verifiable
  Credentials and Zero-Knowledge Proofs](https://arxiv.org/pdf/2209.00586.pdf)
- [BBS Sign](https://asecuritysite.com/golang/bbs_sign)
- [BBS Signature Scheme](https://identity.foundation/bbs-signature/draft-irtf-cfrg-bbs-signatures.html)
- [Arduino API Cloud](https://www.arduino.cc/reference/en/iot/api/)

### Version
*0.1.0*

## Overall Description
We are a data dependent society. Humans rely on data in our daily lives more than ever. We use data for everything from making decisions, staying informed, and connecting with others. Whether it's checking the weather, using GPS for directions, or scrolling through social media, data plays a crucial role in our modern society, shaping how we live and interact with the world around us.

In our society, it is a common belief that data ownership is same as privacy but it is not. Just beacuse you own the data does not mean you get to keep it private. Yet, privacy is still your right.

We also have frequent data breaches. In last few years, we can hear people saying that decentralization is the privacy but it is not.Decentralization is about not relying on one central authority, instead it is more like spreading the responsibilities across different entities.

### Product Perspective
This product is supposed to be an open-source. It is a node based system (via web interface) that allows users to participate in the Averveil ecosystem.

### Product Features
Node: A node based system that allows users to participate in the Averveil ecosystem by connecting to it.
Escrow: A service that allows transaction amongst 2 parties with the help of a third party where third party acts as keeper of the funds.
ZK-Insights: Service that collects the information from nodes and other services to create insights on them.
ZK-Marketing: Service that allows companies to show ads to nodes without having any knowledge of them
ZK-Report: TBD
DAO: Responsible for making decisions on behalf of the system.

### User Classes and Characteristics
Node operators:
- Help the ecosystem run.
- Provide with data from various data sources.
- Are active users who engage with system to make use of different services.

Business Users:
- They can have multiple goals e.g. they might want to serve ads to users.
- Responsible for the direct payments to the ecosystem and are considered bulk or big buyers of native token.

## MVP Focus
### Problem Statement
When you use services, you often share your data willingly or unknowingly, like when visiting websites or using apps. While you may think your data is kept private, recent data breaches show it's not always the case; your information isn't as secure as you might believe.

Similarly, thinking you're private just because you're decentralized isn't accurate. Decentralization means being more independent and responsible for your data, but it doesn't guarantee complete privacy.

### Solution Hypothesis
**Problem:** In today's world, data is incredibly important, but it's getting more complex and valuable. However, we often lack control over our data, and distinguishing real from fake data is difficult.

**Solution:** We need a system that lets us decide who we share our data with, take responsibility for our choices, and even receive incentives for sharing. Additionally, if we choose to share it against some incentive, someone else should be able to verify the source of the data to provide some accountability.

**Result:** With this approach, we can gain more control over our data, ensure its accuracy, and potentially benefit from sharing it while protecting our privacy and security.

### MVP Definition
A basic version of the data control and verification system that allows users to choose data sharing preferences, take responsibility for their data, receive incentives for sharing, and verify data accuracy. This MVP aims to provide essential functionality while serving as a foundation for future enhancements and features. Additionally, entities with whom data is shared can benefit from using the data, being able to verify the source and accuracy of data without actually being able to see the data.

### Build-Measure-Learn-Loop
1. **Build:** Start by developing the most basic version of your product or system that includes the core features necessary to solve the identified problem. This should be the MVP you defined.

2. **Measure:** Once your MVP is built, release it to a small group of users or your target audience. Collect relevant data and feedback on how they interact with the MVP. This could include usage statistics, user feedback, and any other relevant metrics.

3. **Learn:** Analyze the data and feedback you've gathered from the users. Pay attention to what's working well and what needs improvement. Identify any unexpected challenges or opportunities that arise from user interactions with the MVP.

4. **Iterate:** Based on the insights and lessons learned, make necessary adjustments and improvements to your MVP. This could involve adding new features, refining existing ones, or addressing issues that users encountered.

5. **Repeat:** Continue the cycle by releasing the updated MVP to a new group of users or the same group if appropriate. Repeat the process of measuring and learning, making incremental improvements each time.

By using the MVP in the Build-Measure-Learn loop, we can gradually refine our product, ensuring it aligns more closely with user needs and preferences. This iterative approach helps us to avoid spending excessive time and resources on features that may not be valuable or necessary, ultimately increasing the chances of creating a successful and user-centered product.

## System Features and Requirements
<!-- ### Feature 1
#### Description
##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria -->

### Node
#### Description
Node is a decentralized node-based system designed to facilitate the seamless operation of various services within the Averveil multiverse as well as other future multiverses.
Node enables users to create their own user profiles and show proof of these profiles. This helps the Averveil ecosystem gather more information about users, allowing other entities to benefit from these insights. Nodes are rewarded with incentives for sharing this data. Nodes need to provide a verifiable proof which entities can verify.
Averveil serves as a medium for entities and users in their communication. Nodes also have the responsibility of signing their data, which allows for tracing back to the source and ensures privacy.
Nodes can link various sensors through an Arduino-based client. This data can be gathered, converted, and shared using verifible credentials and zero knowledge proofs with other entities. Averveil acts as the middleman for communication, verification, and validation. This approach draws inspiration from the ideas presented in this [paper](https://arxiv.org/pdf/2209.00586.pdf).

##### Priority
:star::star::star::star::star:

#### Functional Requirements
- Authentication and Authorization
- Access Control
- Blockchain Interface
- Storage Interface
   - Local Storage
   - Swarm
- Cryptography
- ZKP
- Signing Interface
- Sensors Interface
   - Transcoder
   - Reciever
   - Signer
   - Prover
- User Interface
#### User Stories
- As a user, I want to be able to control who can access my data and to what extent would this data be accessible.
- I want to be able to prove to someone that some data `D` belongs to me or has been generated by me.
- I want to be able to prove that the data I generated is correct and is signed by me.
- I want to be incentivized for usage of my data.
- I want to be able to connect different sensors with my node.
- I want to be able to generate zero knowledge proofs about my data for validation purposes.
#### Validation Criteria
- I want to be able to own verificable credentials that link my data to me.
- I want assurances that my data will be kept private and secure.
- I want assurances that I will be incentivised.
- I want assurances that only I can allow access to my data.

### Escrow
#### Description
A blockchain based temporary legal arrangement between 2 transacting parties where a third party holds the financial payment.This innovative solution aims to provide a secure and transparent platform for facilitating transactions. By harnessing the power of blockchain technology, the app ensures that transactions are executed only when predetermined conditions are met, eliminating the need for traditional intermediaries and fostering a high level of trust among participants. This project aligns with the ethos of decentralization and self-executing smart contracts, contributing to a more efficient and equitable digital economy.

##### Priority
:star::star::star:
#### Functional Requirements
- Authorization and verification
- User facing integration
- Escrow service
- Payments and commissions management
- Conditional execution of transations

#### User Stories
- I want to be able to send money to someone and they should be able to recieve the money.
- The money should be recieved only if certain per-defined conditions are met.
- If condtions are not met, I want my money returned to me.

#### Validation Criteria
- User must be able to send and receive money.
- User must be able to define conditions for exchange of hands.

### ZK-Insights
#### Description

##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria

### ZK-Marketing
#### Description
##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria

### ZK-Reports
#### Description
##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria

### ERC20 Token
#### Description
##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria

### DAO
#### Description
##### Priority
#### Functional Requirements
#### User Stories
#### Validation Criteria

## Interface Requirements
### User Interfaces
- Website
- Backend application
- Mobile Application
### Software Interfaces
- API
- MQTT Brokers
- Docker
- Blockchain (TBD) and Contracts
- Arduino
### Hardware Interfaces
- Mobile phone
- Computer or Laptop
- Arduino with sensors

## Non-Functional Requirements
### Performance
### Scalability
### Usability
### Security

<!-- ## Lean Metrics
### Engagement
### Growth
### Monetization
### Marketing
 -->

## Lean Metrics
### Engagement:

**Engagement Strategy:**
Averveil plans to harness various online platforms to foster user interaction and brand commitment. Utilizing Discord for daily AMAs and LinkedIn for regular posts, Averveil aims to establish an informative dialogue with its audience. Blogging and newsletters will serve as educational tools, enlightening the community about Web3 based solutions.

**KPIs for Engagement:**
- Click-Through Rate (CTR)
- Conversion Rate
- Impressions
- Bounce Rate
- Ad Engagement
- Mobile Performance

### Growth:

**Growth Strategy:**
To expand its user base and online presence, Averveil will employ SEO strategies, back-linking, and account growth plans on LinkedIn. The aim is to redirect traffic to Averveil’s main website and enhance online visibility. Paid advertisements on LinkedIn, Facebook, and Google will also play a pivotal role in driving targeted traffic and expanding awareness.

**KPIs for Growth:**
- Ad Position
- Quality Score
- Geographic Performance
- Ad Schedule Performance
- Brand Awareness

### Monetization:

**Monetization Strategy:**
Averveil offers distinct products, each contributing to revenue generation. The pricing strategy is under development and will be communicated transparently as the project matures. The focus will be on providing value while ensuring the financial sustainability of the platform.

**KPIs for Monetization:**
- Return on Investment (ROI)
- Cost Per Click (CPC)
- Customer Acquisition Cost (CAC)
- Customer Lifetime Value (CLV)

### Marketing:

**Marketing Strategy:**
Averveil’s marketing approach is multi-faceted, encompassing content sharing through LinkedIn, Discord, and blogs. The aim is to raise awareness, educate, and build trust within the digital ecosystem. Paid advertising on social media platforms and Google will enhance brand visibility and lead generation, with performance monitored through various KPIs.

**KPIs for Marketing:**
- Conversion Value
- Ad Copy and Design Testing
- Ad Position Share
- Ad Extensions
- Impressions

## Appendix
### Assumptions and Dependencies
### Future Iterations

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
- [ ] Set up and configure Swarm and libp2p for data storage.

##### Milestone Q2 and Q3: Contract and Services
- [ ] Design the ERC20 Contract.
- [ ] Define tokenomics and DAO structure.
- [ ] Design the structure of liquidity pools.
- [ ] Identify and implement mock contracts for all services.
- [ ] Implement smart contracts for liquidity pool management.
- [ ] Implement the Mock Services idenfitifed in Sprints with value objects and data transfer objects.
- [ ] Re-configure the node structure based on the newly identified mocks of other services with appropriate data transfer objects and value objects.

##### Milestone Q4: ZKP and Sensors
- [ ] Implement ZKP and Sensors implementation.
- [ ] Averveil frontend implementation.
- [ ] Averveil MVP Completion.
- [ ] Data-remodelling for storage on Swarm.