# Copilot Instructions for the `go-learn` hands-on course

Purpose
-------
This repository is a set of lessons and exercises for learning Go. Use Copilot
to help complete exercises, generate small helper functions, and suggest
explanations. Keep suggestions focused on short, testable code snippets.

How to help a learner (guidelines for Copilot)
---------------------------------------------
- Prefer simple, idiomatic Go. Keep functions small (<= 25 lines).
- When adding example code, include a one-line comment describing intent.
- When asked to implement an exercise, produce both:
  1) A minimal, correct implementation.
  2) A short explanation (1-3 sentences) of why it works and edge-cases.
- If exercise requires IO, prefer functions returning values so they are easy
  to test and run from the course runner.
- Add unit-test style examples where practical (small table-driven tests).

Repository conventions
----------------------
- Lessons live in top-level folders like `01-easy`, `02-intermediate`, `03-advanced`.
- This course also contains a `course` package providing a runnable registry
  of exercises and a runner at `cmd/course`.
- Exercise function naming: LessonShort_ExN (e.g., Hello_Ex1).

When you modify files
---------------------
- Add small, focused edits. Avoid reformatting unrelated code.
- Ensure `go build ./...` succeeds after changes.
- If you add exported identifiers, document them with a short comment.

If uncertain, ask the user one clarifying question rather than guessing.

-- End of Copilot instructions
