🚀 Workshop Demo: Multi-Agent Orchestration in Action
Scenario: The user wants to add a "Newsletter Signup" feature to an existing website. Instead of writing the code, the user acts as the Director, orchestrating three specialized agents.

🎭 The Cast of Agents
Lead Architect: Analyzes the codebase and creates the "Master Plan."

The Developer: Implements the logic and UI based on the plan.

The QA/Security Guard: Reviews code for bugs and security flaws (like SQL injection).

🏁 Phase 1: The "Handoff" from Human to Architect
Action: Open your terminal/IDE and send the high-level goal to the Lead Architect.

User Prompt: > "I need to add a newsletter signup footer. Use the Architect agent to scan the repo and create a PLAN.md. Don't write any code yet. Just tell me how you're going to integrate this with our existing Supabase backend."

What to point out to the team:

Notice how the agent doesn't start coding immediately.

It is "Grounding" itself by reading the schema.sql and Folder Structure.

Result: The agent generates a PLAN.md file.

⛓️ Phase 2: Sequential Handoff (Architect → Developer)
Action: Review the plan, then trigger the Developer.

User Prompt: > "The plan looks solid. Developer, take the PLAN.md and execute the implementation. Build the React component and the API route. Ensure you use the Tailwind variables defined in theme.js."

What to point out to the team:

The Shared Context: The Developer isn't guessing; it is reading the "Artifact" (the PLAN.md) created by the previous agent.

A2A Interaction: If the Developer finds a contradiction in the plan, it will "call back" the Architect (or ask you) for clarification.

Result: New files appear in the sidebar (Newsletter.tsx, api/subscribe.ts).

🛡️ Phase 3: The "Mesh" Interaction (Developer + Security)
Action: Now, show how a secondary agent "watches" the work. This simulates the Security MCP or a specialized Reviewer Agent.

User Prompt: > "Security Agent, review the code just written by the Developer. Specifically check the API route for input validation and rate limiting."

The "Magic" Moment (Live Correction):

Security Agent: "I've found a vulnerability in api/subscribe.ts. The email input isn't being sanitized before the database query. I am providing a fix to the Developer now."

Developer Agent: "Fix received. Updating the API route with Zod validation."

What to point out to the team:

This is Agent-to-Agent correction. \* The user didn't have to find the bug; the Security Agent acted as a "Continuous Auditor."

🚦 Phase 4: Final Validation & The "Human-in-the-loop"
Action: The agents have finished. Now you demonstrate the Verification Skill.

User Prompt: > "Run the local dev server, find the component on the page, and simulate a successful signup. If it works, provide a summary of changes."

Result: \* The agent uses a Browser Tool (or Terminal) to verify the UI.

It confirms: "Signup successful. Entry confirmed in Supabase table 'subscribers'."

💡 Key Takeaways for the Team
Spec-Driven: We started with a plan, not code.

Delegation: We didn't "prompt" the code; we managed a workflow.

Orchestration: Each agent had a specific "Identity" (Architect, Coder, Security).

Efficiency: The total "Human Effort" was 4 prompts. The "AI Effort" was complex file analysis, coding, and security auditing.
