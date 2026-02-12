# 00 -- System Design Answer Framework

A repeatable structure for answering any system design question in an interview.

---

## The 7-Step System Design Answer

1. **Clarify requirements** -- ask 2-3 questions (scope, users, scale)
2. **Define functional requirements** -- what the system must do (APIs, core features)
3. **Define non-functional requirements** -- latency, throughput, availability, durability
4. **Estimate scale** -- requests/sec, storage size, bandwidth (back-of-envelope)
5. **High-level design** -- draw boxes: client, LB, services, DB, cache, queue
6. **Deep-dive one component** -- pick the hardest part and design it thoroughly
7. **Discuss trade-offs and bottlenecks** -- what breaks at 10x scale, what you would change

---

## What Interviewers Actually Test

- **Structured thinking** -- can you break a vague problem into clear steps?
- **Breadth** -- do you know the major building blocks (cache, queue, LB, DB)?
- **Depth** -- can you deep-dive into one component with real detail?
- **Trade-off awareness** -- do you weigh options or just pick one?
- **Communication** -- can you explain clearly while drawing on a whiteboard?
- **Prioritization** -- do you focus on what matters or get lost in details?
- **Handling unknowns** -- do you ask questions or assume?

---

## How to Speak Confidently

- Start every answer with: "Let me first clarify the requirements..."
- Use phrases like: "The trade-off here is..." / "One approach is X, but Y gives us..."
- When stuck: "Let me think about this for a moment" (buy time, do not panic)
- Name patterns explicitly: "I would use a **token bucket** for rate limiting because..."
- Mention monitoring early: "We would need metrics on latency and error rates"
- End with: "If we needed to scale further, I would consider..."
- Never say "I do not know" without adding "but I would approach it by..."

---

## TL;DR

- Always follow a structure -- interviewers notice when you wing it
- Clarify before designing -- wrong assumptions waste 10 minutes
- Non-functional requirements separate junior from senior answers
- Deep-dive shows real engineering skill, not just buzzword knowledge
- Trade-offs are the most important part -- always discuss at least two options
- Practice the 7 steps on every problem until it becomes automatic
