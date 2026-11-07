import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './style.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <main><h1>Scam Incident Registry</h1><p>An open-source Nigerian scam incident reporting and lookup system.</p></main>
  </StrictMode>,
);
