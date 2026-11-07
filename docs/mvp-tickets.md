# SIR MVP ticket drafts

Source: Product Brief & MVP Scope v1.0, updated 6 November 2026.

These are draft tickets, not issues created in a tracker. IDs are local references. Delivery order is suggested; estimates and assignees are intentionally unset. Every ticket is part of MVP unless explicitly identified as a decision. The implemented baseline uses Go standard-library HTTP routing and React with TypeScript and Vite; remaining architecture and operational decisions are tracked in SIR-01. Platform usernames are an extension point, not an MVP search type.

## Tracking

Use **Planned → In progress → Done**. Mark a ticket Done only when all its
acceptance criteria are implemented and verified. Checkboxes describe complete
criteria; progress notes record partial delivery. These statuses describe the
repository, not external tracker issues.

## Foundation

### SIR-01 — Record architecture and operational decisions

**Status:** In progress

**Outcome:** Contributors have an agreed implementation baseline.

**Acceptance criteria**
- [ ] Record decisions for the frontend, router, relational database, database access, storage, background jobs, API documentation, deployment, and observability; distinguish proposed from approved choices.
- [ ] Define module boundaries for authentication, incidents, identifiers, evidence, moderation, disputes, search, and API clients.
- [ ] Document pending operational decisions: account verification method, identifier publication rules, file limits, usage limits, and moderation policies.
- [ ] Keep unapproved recommendations and the illustrative 12-week schedule out of delivery commitments.

**Progress:** Monorepo, React/TypeScript frontend, standard-library Go router,
domain boundaries, and OpenAPI location are established. Shared infrastructure
is grouped under internal/platform; feature handlers and queries belong to domains. See MEMORY.md for
confirmed decisions and outstanding choices. The broader criteria remain open.

**Dependencies:** None.

### SIR-02 — Bootstrap the Go service and contributor workflow

**Status:** In progress

**Outcome:** Developers can run and validate the backend locally.

**Acceptance criteria**
- [ ] A Go service starts using documented configuration and exposes health/readiness checks.
- [ ] Establish module structure, request validation, consistent errors, and structured logging without credentials or sensitive payloads.
- [ ] Provide local setup instructions and example configuration without secrets or real incident data.
- [ ] CI runs formatting, static checks, and automated tests.

**Implemented foundation**
- [x] Runnable API with GET/HEAD health endpoint and OpenAPI documentation.
- [x] Validated PORT configuration, structured lifecycle logs, HTTP timeouts, and graceful shutdown.
- [x] Domain directory structure and documented local commands.
- [x] Configuration, routing, and shutdown tests; race tests, vet, build, and health/SIGTERM smoke test passed during implementation.

- [x] Opt-in Go Postgres/S3 client integration tests and safe local example configuration.
- [x] Infrastructure Make targets read `backend/.env` for the LocalStack token; Compose configuration validated with a synthetic env-file token, and local `.env` confirmed git-ignored.
- [x] Integration Make target loads `backend/.env`; env-file exports/precedence, environment-only mode, and missing-file failure verified with a synthetic command harness. Backend race tests, vet, formatting, and build pass.
- [ ] Live connection checks verified against local infrastructure (Docker unavailable during implementation).

Integration setup follow-up: missing-variable errors are resolved. Live Postgres
and S3 checks still fail at connection checks; the Docker daemon is unreachable
from the verification environment, including outside the sandbox.

**Remaining:** Readiness semantics, request-body validation and consistent API
errors for domain endpoints and CI. The bundled
acceptance criteria above remain unchecked until fully satisfied.

**Dependencies:** SIR-01.

### SIR-03 — Reconcile the schema and add database migrations

**Status:** Planned

**Outcome:** The database supports the finalised MVP scope.

**Acceptance criteria**
- [ ] Update `schema.mermaid` and create migrations covering users/RBAC, incidents, canonical identifiers and links, evidence, moderation, API clients, summaries, disputes, and audit events.
- [ ] Add category, approximate occurrence dates, optional payment channel, loss disclosure consent, and private reporter contact support missing from the existing draft.
- [ ] Model report revisions, public redacted content, evidence extracts, and lifecycle transitions without overwriting review history.
- [ ] Enforce foreign keys, canonical identifier uniqueness, and appropriate indexes and constraints.
- [ ] Verify migrations against an empty database; document migration recovery procedures and provide synthetic development data.

**Dependencies:** SIR-01, SIR-02.

### SIR-04 — Implement reporter registration, verification, and sign-in

**Status:** Planned

**Outcome:** Registered, verified individuals can access reporting features.

**Acceptance criteria**
- [ ] Users can register, complete the agreed verification process, sign in, sign out, and recover account access.
- [ ] Verification and recovery tokens expire and cannot be reused; passwords and sessions use secure handling.
- [ ] Anonymous, unverified, and inactive users cannot submit reports.
- [ ] Reporter identity and contact information remain private.
- [ ] Authentication and recovery endpoints have abuse protections.

**Dependencies:** SIR-01, SIR-03.

### SIR-05 — Enforce role and ownership permissions

**Status:** Planned

**Outcome:** Reporters, moderators, administrators, and API clients have appropriate access.

**Acceptance criteria**
- [ ] Define and implement a permission matrix with default-deny behaviour.
- [ ] Reporters can access their own private submissions; accessing another reporter's private records is denied.
- [ ] Moderation, evidence access, API-client management, and role assignment require explicit permissions.
- [ ] Public and API-client access cannot expose private reporter or evidence data.
- [ ] Automated tests cover ownership bypasses and privilege escalation attempts.

**Dependencies:** SIR-04.

### SIR-06 — Implement restricted audit history

**Status:** Planned

**Outcome:** Sensitive and administrative actions are attributable and reviewable.

**Acceptance criteria**
- [ ] Record actor, action, target, timestamp, and relevant decision/reason for moderation, evidence access, disputes, permission changes, and API credential administration.
- [ ] Audit events cannot be edited through ordinary application workflows.
- [ ] Access is restricted; logs exclude passwords, credentials, evidence contents, and unnecessary personal data.
- [ ] Retention and deletion behaviour follow the approved policy from SIR-21.

**Dependencies:** SIR-03, SIR-05; retention configuration depends on SIR-21.

## Reporting and evidence

### SIR-07 — Normalise and deduplicate supported identifiers

**Status:** Planned

**Outcome:** Equivalent identifiers link to one canonical record across incidents.

**Acceptance criteria**
- [ ] Support Nigerian phone numbers in E.164, bank code plus account number, email addresses, and URLs.
- [ ] Define and document type-specific normalisation rules, preserving meaningful distinctions and bank-account leading zeros.
- [ ] Keep submitted raw values separate from canonical values with appropriate access restrictions.
- [ ] Concurrent submissions of the same canonical identifier do not create duplicates.
- [ ] Test equivalent inputs, invalid inputs, and distinct identifiers that must not be merged.
- [ ] Do not implement alleged-party name or platform-username search.

**Dependencies:** SIR-03.

### SIR-08 — Implement private incident drafts and submission

**Status:** Planned

**Outcome:** Verified reporters can document incidents they personally experienced.

**Acceptance criteria**
- [ ] Reporters can create and edit private drafts with title, description, category, occurrence date/precision, and at least one supported identifier before submission.
- [ ] Support optional amount/currency, disclosure consent, payment channel, alleged-party details, and evidence references.
- [ ] Capture confirmation that the reporter personally experienced the incident.
- [ ] Validate required fields and identifier links server-side; evidence is encouraged but optional.
- [ ] Submission moves the report from Draft to Submitted without making it publicly accessible.
- [ ] Reject unauthorised lifecycle transitions and changes to another user's report.

**Dependencies:** SIR-05, SIR-07.

### SIR-09 — Build the reporter form and submission dashboard

**Status:** Planned

**Outcome:** Reporters can submit reports and follow their progress.

**Acceptance criteria**
- [ ] Provide an accessible reporting form with draft saving, approximate date input, field validation, identifier entry, and optional evidence upload.
- [ ] Explain private fields, publication review, and optional loss disclosure before submission.
- [ ] Show the reporter's drafts, submitted reports, current status, and moderator requests for changes.
- [ ] Allow requested revisions and resubmission while preserving review history; revisions to published reports require review before public replacement.
- [ ] Cover empty, loading, validation, and submission-error states.

**Dependencies:** SIR-04, SIR-08, SIR-10; moderator feedback integrates with SIR-12.

### SIR-10 — Add private evidence uploads and malware screening

**Status:** Planned

**Outcome:** Evidence can be submitted without exposing unsafe or private files.

**Acceptance criteria**
- [ ] Authorised reporters can attach files to their own incidents using private storage.
- [ ] Enforce configured size and type limits, validate file content, calculate integrity hashes, and record uploader metadata.
- [ ] Quarantine files until scanning succeeds; infected, failed, or pending files cannot be served as safe evidence.
- [ ] Evidence retrieval requires authorisation and uses controlled access; publishing an incident never changes original evidence visibility.
- [ ] Audit evidence access and prevent cross-user retrieval or guessing of storage references.

**Dependencies:** SIR-01, SIR-05, SIR-06, SIR-08.

### SIR-11 — Support separately reviewed evidence extracts

**Status:** Planned

**Outcome:** Moderators can publish selected redacted extracts while originals remain private.

**Acceptance criteria**
- [ ] Authorised moderators can create or upload a redacted derivative associated with an original evidence item.
- [ ] Extracts receive validation, scanning, explicit review, and publication approval before public access.
- [ ] Public responses contain only approved extracts, never original file references or sensitive metadata.
- [ ] Revoking extract approval or unpublishing the incident removes public access, including cached or temporary access within a documented bound.
- [ ] Record extract creation, approval, and revocation in audit history.

**Dependencies:** SIR-10, SIR-12, SIR-21.

## Moderation and accountability

### SIR-12 — Implement moderation decisions and publication rules

**Status:** Planned

**Outcome:** Only reviewed, appropriately redacted reports become public.

**Acceptance criteria**
- [ ] Support Draft, Submitted, Under Review, Published, and Rejected states with authorised transitions.
- [ ] Moderators can request changes, redact public content, approve, reject, correct, and unpublish reports with recorded reasons.
- [ ] Preserve original restricted submissions and previous decisions/revisions.
- [ ] Track disputes, corrections, and withdrawal without losing the report's history.
- [ ] Publish only explicitly approved fields, date granularity, identifiers, and consented/reviewed loss information.
- [ ] Prevent concurrent reviews from silently overwriting each other or publishing an unreviewed revision.
- [ ] Publication wording explains that moderation is not independent proof of an allegation.

**Dependencies:** SIR-06, SIR-08, SIR-21 publication rules.

### SIR-13 — Build the moderator review dashboard

**Status:** Planned

**Outcome:** Moderators can process submissions and inspect their context.

**Acceptance criteria**
- [ ] Provide queues filtered by status with incident details, linked identifiers, restricted alleged-party information, and authorised evidence access.
- [ ] Present original content and a preview of the public version distinctly.
- [ ] Provide review actions, change requests, redaction controls, and decision history.
- [ ] Require reasons where defined by moderation policy and surface stale-review conflicts.
- [ ] Enforce permission checks on both dashboard routes and backend actions.

**Dependencies:** SIR-10, SIR-12.

### SIR-14 — Add disputes, correction requests, and withdrawal handling

**Status:** Planned

**Outcome:** Affected parties and reporters can request review of a published report.

**Acceptance criteria**
- [ ] Provide an accessible way for affected parties to submit a dispute or correction request, including report reference, reason, and private contact information.
- [ ] Reporters can request correction or withdrawal of their own reports.
- [ ] Apply input validation and abuse protections; define requester verification in operational policy without imposing the reporter-only submission rule on affected parties.
- [ ] Moderators can track, review, and resolve requests with reasons and communicate the outcome through the agreed channel.
- [ ] Serious disputes can trigger temporary unpublication during review; resolved corrections require review before republication.
- [ ] Dispute contacts and supporting material remain private, and decisions are audited.

**Dependencies:** SIR-06, SIR-12, SIR-21.

## Public discovery and API

### SIR-15 — Implement shared public-safe queries and identifier summaries

**Status:** Planned

**Outcome:** Web search and the external API return consistent, publishable information.

**Acceptance criteria**
- [ ] Use the same canonicalisation as reporting for phone, bank account, email, and URL lookup.
- [ ] Return only published, approved incident projections with counts, categories, approximate dates, and detail references.
- [ ] Exclude private reporter data, restricted alleged-party details, unredacted narratives, and original evidence references.
- [ ] Update summaries on publication, correction, identifier reassignment, and withdrawal; public reads exclude unpublished records even if aggregate refresh is delayed.
- [ ] Define and test aggregate freshness and pagination; provide a summary rebuild/reconciliation mechanism.
- [ ] Do not add name search, bulk exports, guilt labels, or fraud scores.

**Dependencies:** SIR-07, SIR-12.

### SIR-16 — Build public identifier search and incident pages

**Status:** Planned

**Outcome:** Visitors can look up identifiers without registration and read moderated reports.

**Acceptance criteria**
- [ ] Provide phone, bank account, email, and URL search with type-specific validation and bank selection/code input where needed.
- [ ] Results show published counts, categories, approximate dates, and links to approved incident details.
- [ ] Detail pages display only the public-safe projection and explicitly approved evidence extracts.
- [ ] Explain that identifier matches do not prove wrongdoing and may involve reassigned, shared, or impersonated identifiers; no-result copy does not imply safety.
- [ ] Provide dispute/correction entry points and appropriate unavailable states for unpublished reports.
- [ ] Include accessible loading, empty, invalid-input, and rate-limit states; do not offer name search.

**Dependencies:** SIR-11, SIR-14, SIR-15, SIR-17.

### SIR-17 — Protect public search against abuse and enumeration

**Status:** Planned

**Outcome:** Anonymous access is bounded and misuse is observable.

**Acceptance criteria**
- [ ] Apply configurable request and result limits, bounded pagination, and controls against automated bulk enumeration.
- [ ] Cover both public web endpoints and alternate access paths; reject unsupported search modes.
- [ ] Return consistent rate-limit responses with retry guidance.
- [ ] Record restricted operational signals without exposing raw search identifiers in general logs.
- [ ] Test burst traffic, repeated enumeration attempts, and deployment-wide enforcement using approved pilot thresholds.

**Dependencies:** SIR-01, SIR-06, SIR-15.

### SIR-18 — Implement approved API-client provisioning and credential lifecycle

**Status:** Planned

**Outcome:** Administrators can grant and revoke free API access.

**Acceptance criteria**
- [ ] Administrators can record an organisation, approval decision, acceptable-use acknowledgement, and configured usage limits.
- [ ] Issue credentials only to approved clients; show secrets only at issuance and store a secure verification representation.
- [ ] Support credential rotation, revocation, expiry, and client deactivation with audited actions.
- [ ] Reject missing, invalid, expired, or revoked credentials and enforce per-client limits.
- [ ] Do not add billing, paid tiers, or a self-service API dashboard to MVP.

**Dependencies:** SIR-05, SIR-06, SIR-21 acceptable-use rules.

### SIR-19 — Ship the versioned third-party search API

**Status:** Planned

**Outcome:** Approved clients can retrieve published incident information.

**Acceptance criteria**
- [ ] Implement authenticated `GET /v1/identifiers/search` and `GET /v1/incidents/{id}` using shared public-safe domain queries.
- [ ] Document identifier parameters, pagination, response fields, authentication, errors, usage limits, and examples in an API contract.
- [ ] Enforce client limits and restricted usage logging; unpublished incidents and private fields are never returned.
- [ ] Contract tests cover all four identifier types, no matches, unpublished records, invalid credentials, and rate limits.
- [ ] Access is free; no bulk export or name-search endpoint is exposed.

**Dependencies:** SIR-15, SIR-17, SIR-18.

## Release readiness

### SIR-20 — Add deployment, monitoring, backups, and restore procedures

**Status:** Planned

**Outcome:** Operators can run the pilot and recover its data.

**Acceptance criteria**
- [ ] Document deployment, configuration, secret handling, migrations, and rollback procedures for the agreed infrastructure.
- [ ] Configure transport/storage encryption and approved protection for sensitive fields.
- [ ] Monitor service errors, latency, evidence scanning failures, background-job failures, and abuse signals without leaking personal data.
- [ ] Back up the database and private evidence; complete and record a restore drill against agreed recovery targets.
- [ ] Document access control, alert ownership, and incident response procedures.

**Dependencies:** SIR-01, SIR-06, SIR-10; final drill requires integrated MVP.

### SIR-21 — Finalise publication, privacy, and operational policies

**Status:** Planned

**Outcome:** Implementation and launch have explicit approved rules for handling allegations and personal data.

**Acceptance criteria**
- [ ] Obtain qualified review of data-processing grounds under Nigeria's Data Protection Act 2023 and whether a DPIA is required; record the outcome and complete required actions.
- [ ] Approve identifier exposure/masking, narrative redaction, date granularity, loss consent, evidence extracts, and moderation criteria.
- [ ] Define retention, correction, withdrawal, deletion, backup expiry, and restricted-log handling; assign implementation owners.
- [ ] Publish applicable privacy, acceptable-use, reporting, and dispute guidance with allegation-aware language.
- [ ] Define verification, moderation escalation, dispute handling, and requester communication procedures.
- [ ] Record launch approval and unresolved blockers. This ticket specifies a review deliverable, not a legal conclusion.

**Dependencies:** None to start; validate final policies against the implemented workflows.

### SIR-22 — Enforce retention and deletion policies

**Status:** Planned

**Outcome:** Approved data-handling rules are executable and auditable.

**Acceptance criteria**
- [ ] Implement restricted handling of correction/deletion requests and configured retention jobs for applicable records, evidence, and logs.
- [ ] Withdrawal promptly removes public visibility while retention/deletion follows the approved policy.
- [ ] Remove or expire derivatives, cached public content, and storage access where required; document backup expiry and restoration handling.
- [ ] Ensure aggregates remain consistent and retained audit information does not preserve prohibited personal data.
- [ ] Verify policy behaviour using synthetic expired records and deletion requests.

**Dependencies:** SIR-14, SIR-20, SIR-21.

### SIR-23 — Instrument pilot success metrics

**Status:** Planned

**Outcome:** The pilot measures usefulness and operational quality.

**Acceptance criteria**
- [ ] Define and measure submission completion, moderation turnaround, published coverage, match rate, search-to-detail engagement, disputes/corrections, privacy incidents, API performance, and approved-client adoption.
- [ ] Define how to measure useful moderated matches, including the limits of any proxy metric.
- [ ] Collect only necessary telemetry and exclude raw identifiers, reporter identities, and evidence contents from analytics.
- [ ] Provide a restricted pilot report or dashboard with documented metric definitions.

**Dependencies:** SIR-09, SIR-14, SIR-16, SIR-19, SIR-21.

### SIR-24 — Validate the complete MVP and run a controlled pilot

**Status:** Planned

**Outcome:** Release readiness is evidenced against the brief's acceptance criteria.

**Acceptance criteria**
- [ ] Exercise registration → verification → report submission → evidence scan → moderation → public search → API retrieval using synthetic data.
- [ ] Verify shared canonical identifiers across reports and all four supported search types.
- [ ] Verify corrections, disputes, unpublication, and withdrawal remove inappropriate public access and update summaries.
- [ ] Test unauthorised access, private-file leakage, unreviewed revisions, permission bypasses, abuse controls, and absence of name search.
- [ ] Confirm monitoring, tested recovery, policy implementation, and approved legal/privacy safeguards.
- [ ] Run a bounded pilot with assigned operators, documented stop/rollback criteria, and recorded findings; resolve release blockers before expansion.

**Dependencies:** SIR-01–SIR-23.

## Suggested delivery sequence

1. Start architecture decisions and policy review together: SIR-01, SIR-21.
2. Establish service, schema, identity, permissions, and audit: SIR-02–SIR-06.
3. Deliver reporting and private evidence: SIR-07–SIR-10.
4. Complete moderation, reviewed extracts, and disputes: SIR-12, SIR-13, SIR-11, SIR-14.
5. Deliver shared search, protections, public views, and approved API access: SIR-15, SIR-17, SIR-16, SIR-18, SIR-19.
6. Develop operations alongside features; finish retention, metrics, and pilot validation: SIR-20, SIR-22–SIR-24.

Excluded: AI detection, automated scoring, ML clustering, bank integrations, public comments, native apps, multi-country support, subscriptions, real-time alerts, bulk exports, name search, anonymous incident submission, and automated account freezes.
