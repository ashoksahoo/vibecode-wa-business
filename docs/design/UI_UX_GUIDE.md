# UI/UX Design Guide 🍓
# Vibecoded WA Client

**Design System:** Strawberry Theme  
**Version:** 1.0.0  
**Last Updated:** November 20, 2025

---

## 🎨 Color Palette

### Primary (Strawberry Red)
```javascript
strawberry: {
  50:  '#fff1f2',  // Lightest - backgrounds
  100: '#ffe4e6',  // Very light - cards
  200: '#fecdd3',  // Light - borders
  400: '#fb7185',  // Medium - hover
  500: '#f43f5e',  // PRIMARY - main brand
  600: '#e11d48',  // Dark - active states
  700: '#be123c',  // Darker - pressed
  900: '#881337',  // Darkest - emphasis
}
```

### Secondary (Fresh Green)
```javascript
leaf: {
  100: '#dcfce7',  // Success backgrounds
  500: '#22c55e',  // SUCCESS - delivered status
  700: '#15803d',  // Dark green
}
```

### Neutral
```javascript
neutral: {
  50:  '#fafafa',  // App background
  100: '#f5f5f5',  // Alternate background
  200: '#e5e5e5',  // Borders
  400: '#a3a3a3',  // Placeholder
  500: '#737373',  // Secondary text
  700: '#404040',  // Body text
  900: '#171717',  // Headings
}
```

### Status Colors
```javascript
--success:   #22c55e  (leaf-500)
--error:     #e11d48  (strawberry-600)
--warning:   #f59e0b
--info:      #3b82f6
```

---

## 📝 Typography

```javascript
// Font
font-family: 'Inter', system-ui, sans-serif

// Scale
text-xs:   12px
text-sm:   14px
text-base: 16px
text-lg:   18px
text-xl:   20px
text-2xl:  24px
text-3xl:  30px

// Weights
font-normal:    400
font-medium:    500
font-semibold:  600
font-bold:      700
```

---

## 📐 Spacing & Layout

```javascript
// Spacing (8px grid)
space-1:  4px
space-2:  8px
space-3:  12px
space-4:  16px
space-6:  24px
space-8:  32px
space-12: 48px

// Border Radius
rounded-lg:   8px   // Buttons
rounded-xl:   12px  // Cards
rounded-2xl:  16px  // Message bubbles
rounded-full: 100%  // Avatars
```

---

## 🧩 Core Components

### Button (Primary)
```html
<button class="
  px-4 py-2 
  bg-strawberry-500 
  text-white 
  font-medium 
  rounded-lg 
  hover:bg-strawberry-600 
  active:bg-strawberry-700
  focus:outline-none focus:ring-2 focus:ring-strawberry-500
  transition-colors duration-150
">
  Send Message
</button>
```

### Button (Secondary)
```html
<button class="
  px-4 py-2 
  bg-white 
  text-strawberry-600 
  border-2 border-strawberry-500 
  rounded-lg 
  hover:bg-strawberry-50
">
  Cancel
</button>
```

### Input Field
```html
<input 
  type="text"
  placeholder="+1234567890"
  class="
    px-3 py-2
    border border-neutral-300
    rounded-lg
    focus:outline-none 
    focus:ring-2 
    focus:ring-strawberry-500
    placeholder:text-neutral-400
  "
/>
```

### Search Input
```html
<div class="relative">
  <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-neutral-400" />
  <input 
    type="text"
    placeholder="Search..."
    class="pl-10 pr-4 py-2 w-full border border-neutral-300 rounded-lg"
  />
</div>
```

### Card
```html
<div class="
  bg-white 
  rounded-xl 
  shadow-sm 
  border border-neutral-200
  p-6
  hover:shadow-md
  transition-shadow
">
  Card content
</div>
```

### Contact Card
```html
<div class="
  flex items-center gap-3
  bg-white 
  rounded-lg 
  p-4
  hover:bg-strawberry-50
  cursor-pointer
  transition-colors
">
  <div class="w-12 h-12 rounded-full bg-strawberry-100 flex items-center justify-center">
    <span class="text-strawberry-700 font-semibold">JD</span>
  </div>
  <div class="flex-1">
    <h4 class="font-medium text-neutral-800">John Doe</h4>
    <p class="text-sm text-neutral-500">+1 234 567 890</p>
  </div>
  <div class="px-2 py-0.5 bg-strawberry-500 text-white text-xs rounded-full">3</div>
</div>
```

### Message Bubble (Outbound)
```html
<div class="flex justify-end mb-4">
  <div class="max-w-[70%]">
    <div class="bg-leaf-100 rounded-2xl rounded-tr-sm px-4 py-2 shadow-sm">
      <p class="text-neutral-800">Hello! How can I help you?</p>
    </div>
    <div class="flex items-center gap-2 justify-end mt-1 px-2">
      <span class="text-xs text-neutral-500">10:30 AM</span>
      <CheckCheckIcon class="w-4 h-4 text-leaf-500" />
    </div>
  </div>
</div>
```

### Message Bubble (Inbound)
```html
<div class="flex justify-start mb-4">
  <div class="max-w-[70%]">
    <div class="bg-white rounded-2xl rounded-tl-sm px-4 py-2 shadow-sm">
      <p class="text-neutral-800">I need help with my order</p>
    </div>
    <div class="flex items-center gap-2 mt-1 px-2">
      <span class="text-xs text-neutral-500">10:28 AM</span>
    </div>
  </div>
</div>
```

### Avatar
```html
<!-- With initials -->
<div class="w-12 h-12 rounded-full bg-strawberry-100 flex items-center justify-center">
  <span class="text-strawberry-700 font-semibold">JD</span>
</div>

<!-- With image -->
<div class="w-12 h-12 rounded-full overflow-hidden">
  <img src="/avatar.jpg" alt="User" class="w-full h-full object-cover" />
</div>

<!-- With online status -->
<div class="relative">
  <div class="w-12 h-12 rounded-full bg-strawberry-100">
    <img src="/avatar.jpg" alt="User" class="rounded-full" />
  </div>
  <div class="absolute bottom-0 right-0 w-3 h-3 bg-leaf-500 border-2 border-white rounded-full"></div>
</div>
```

### Badge
```html
<!-- Success -->
<span class="
  inline-flex items-center gap-1
  px-2 py-1
  bg-leaf-100 text-leaf-700
  text-xs font-medium
  rounded-full
">
  <CheckIcon class="w-3 h-3" />
  Delivered
</span>

<!-- Error -->
<span class="
  inline-flex items-center gap-1
  px-2 py-1
  bg-strawberry-100 text-strawberry-700
  text-xs font-medium
  rounded-full
">
  <XIcon class="w-3 h-3" />
  Failed
</span>
```

### Toast Notification
```html
<div class="
  fixed top-4 right-4
  bg-white
  border-l-4 border-leaf-500
  rounded-lg
  shadow-lg
  p-4
  flex items-start gap-3
  max-w-sm
  animate-slideInRight
">
  <div class="w-5 h-5 bg-leaf-100 rounded-full flex items-center justify-center">
    <CheckIcon class="w-3 h-3 text-leaf-600" />
  </div>
  <div class="flex-1">
    <h4 class="text-sm font-semibold text-neutral-800">Message Sent</h4>
    <p class="text-sm text-neutral-600 mt-0.5">Successfully sent to +1 234 567 890</p>
  </div>
  <button class="text-neutral-400 hover:text-neutral-600">
    <XIcon class="w-4 h-4" />
  </button>
</div>
```

### Loading Spinner
```html
<div class="
  w-8 h-8
  border-4 border-neutral-200
  border-t-strawberry-500
  rounded-full
  animate-spin
"></div>
```

---

## 📱 Main Layout: Messages Page

```
┌─────────────────────────────────────────────────────────────┐
│ Top Nav (Logo + User Menu)                                  │
├──────────────┬──────────────────────────────────────────────┤
│              │                                               │
│  Sidebar     │  Conversation List (30%) │ Message Thread (70%)│
│              │                           │                   │
│  - Messages  │  [Search]                │ [Contact Header]  │
│  - Contacts  │                           │                   │
│  - Templates │  Contact 1  [3 unread]   │ Messages...       │
│  - Calls     │  Contact 2               │ ...               │
│  - Analytics │  Contact 3               │ ...               │
│  - Settings  │  ...                     │                   │
│              │                           │ [Message Input]   │
└──────────────┴───────────────────────────────────────────────┘
```

### Implementation
```html
<div class="h-screen flex flex-col">
  <!-- Top Nav -->
  <nav class="h-16 bg-white border-b border-neutral-200 px-6 flex items-center justify-between">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 bg-strawberry-500 rounded-lg"></div>
      <h1 class="text-xl font-bold">Vibecoded WA</h1>
    </div>
    <div>User Menu</div>
  </nav>
  
  <div class="flex-1 flex overflow-hidden">
    <!-- Sidebar -->
    <aside class="w-64 bg-white border-r border-neutral-200">
      <nav class="p-4 space-y-1">
        <!-- Active -->
        <a href="/messages" class="flex items-center gap-3 px-4 py-3 bg-strawberry-50 text-strawberry-700 font-medium rounded-lg">
          <MessageIcon class="w-5 h-5" />
          <span>Messages</span>
        </a>
        <!-- Inactive -->
        <a href="/contacts" class="flex items-center gap-3 px-4 py-3 text-neutral-600 hover:bg-neutral-50 rounded-lg">
          <ContactsIcon class="w-5 h-5" />
          <span>Contacts</span>
        </a>
      </nav>
    </aside>
    
    <!-- Main Content -->
    <main class="flex-1 flex">
      <!-- Conversation List (30%) -->
      <div class="w-[30%] border-r border-neutral-200 bg-white flex flex-col">
        <div class="p-4 border-b">
          <input type="text" placeholder="Search..." class="w-full px-3 py-2 border rounded-lg" />
        </div>
        <div class="flex-1 overflow-y-auto">
          <!-- Contact cards -->
        </div>
      </div>
      
      <!-- Message Thread (70%) -->
      <div class="flex-1 flex flex-col bg-neutral-50">
        <!-- Header -->
        <div class="h-16 bg-white border-b px-6 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-strawberry-100"></div>
            <div>
              <h2 class="font-semibold">John Doe</h2>
              <p class="text-sm text-neutral-500">+1 234 567 890</p>
            </div>
          </div>
        </div>
        
        <!-- Messages -->
        <div class="flex-1 overflow-y-auto p-6">
          <!-- Message bubbles -->
        </div>
        
        <!-- Input -->
        <div class="bg-white border-t p-4">
          <div class="flex items-end gap-3">
            <textarea rows="1" placeholder="Type a message..." class="flex-1 px-3 py-2 border rounded-lg resize-none"></textarea>
            <button class="bg-strawberry-500 text-white p-3 rounded-lg">
              <SendIcon class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>
    </main>
  </div>
</div>
```

---

## 📱 Responsive (Mobile)

### Breakpoints
```javascript
sm:  640px
md:  768px
lg:  1024px
xl:  1280px
```

### Mobile Layout
- Stack conversation list OR message thread (not both)
- Bottom navigation bar (5 icons)
- 44px minimum tap targets
- Full-width inputs

```html
<!-- Bottom Nav (Mobile only) -->
<nav class="fixed bottom-0 left-0 right-0 bg-white border-t md:hidden">
  <div class="flex justify-around py-2">
    <a href="/messages" class="flex flex-col items-center gap-1 p-2 text-strawberry-600">
      <MessageIcon class="w-6 h-6" />
      <span class="text-xs">Messages</span>
    </a>
    <a href="/contacts" class="flex flex-col items-center gap-1 p-2 text-neutral-600">
      <ContactsIcon class="w-6 h-6" />
      <span class="text-xs">Contacts</span>
    </a>
  </div>
</nav>
```

---

## ✨ Animations

```css
/* Fade In */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

/* Slide Up */
@keyframes slideUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

/* Slide In Right (for toasts) */
@keyframes slideInRight {
  from { opacity: 0; transform: translateX(20px); }
  to { opacity: 1; transform: translateX(0); }
}

/* Usage */
.animate-fadeIn { animation: fadeIn 200ms ease-in; }
.animate-slideUp { animation: slideUp 300ms ease-out; }
.animate-slideInRight { animation: slideInRight 300ms ease-out; }
```

---

## ♿ Accessibility

### Color Contrast (WCAG AA)
✅ All color combinations meet WCAG AA standards
- Text on white: 4.5:1 minimum
- Large text: 3:1 minimum

### Focus States
```css
/* Always include focus states */
focus:outline-none focus:ring-2 focus:ring-strawberry-500 focus:ring-offset-2
```

### ARIA Labels
```html
<button aria-label="Send message">
  <SendIcon />
</button>

<div aria-live="polite">Message sent</div>
```

### Keyboard Navigation
- All interactive elements accessible via Tab
- Enter/Space to activate buttons
- Escape to close modals

---

## 🛠️ Tailwind Config

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    extend: {
      colors: {
        strawberry: {
          50: '#fff1f2',
          100: '#ffe4e6',
          200: '#fecdd3',
          400: '#fb7185',
          500: '#f43f5e',
          600: '#e11d48',
          700: '#be123c',
          900: '#881337',
        },
        leaf: {
          100: '#dcfce7',
          500: '#22c55e',
          700: '#15803d',
        },
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
      },
      animation: {
        fadeIn: 'fadeIn 200ms ease-in',
        slideUp: 'slideUp 300ms ease-out',
        slideInRight: 'slideInRight 300ms ease-out',
      },
      keyframes: {
        fadeIn: {
          from: { opacity: '0' },
          to: { opacity: '1' },
        },
        slideUp: {
          from: { opacity: '0', transform: 'translateY(20px)' },
          to: { opacity: '1', transform: 'translateY(0)' },
        },
        slideInRight: {
          from: { opacity: '0', transform: 'translateX(20px)' },
          to: { opacity: '1', transform: 'translateX(0)' },
        },
      },
    },
  },
}
```

---

## 🎯 Quick Reference

### Common Patterns
```html
<!-- Primary Action -->
<button class="px-4 py-2 bg-strawberry-500 text-white rounded-lg hover:bg-strawberry-600">
  Action
</button>

<!-- Secondary Action -->
<button class="px-4 py-2 border-2 border-strawberry-500 text-strawberry-600 rounded-lg hover:bg-strawberry-50">
  Cancel
</button>

<!-- Success State -->
<div class="bg-leaf-100 text-leaf-700 p-4 rounded-lg">
  Success message
</div>

<!-- Error State -->
<div class="bg-strawberry-100 text-strawberry-700 p-4 rounded-lg">
  Error message
</div>

<!-- Card -->
<div class="bg-white rounded-xl shadow-sm border border-neutral-200 p-6">
  Content
</div>

<!-- Input with Error -->
<input class="border border-strawberry-600 rounded-lg px-3 py-2" />
<p class="text-xs text-strawberry-600 mt-1">Error message</p>
```

---

## 📦 Icon Library

**Recommended:** Lucide React
```bash
npm install lucide-react
```

```jsx
import { Send, Phone, MessageSquare, User } from 'lucide-react';

<Send className="w-5 h-5" />
```

---

## ✅ Design Checklist

Before shipping:
- [ ] Uses strawberry/leaf color palette
- [ ] Proper spacing (8px grid)
- [ ] Has hover/focus/active states
- [ ] Includes loading states
- [ ] Shows error states
- [ ] Works on mobile
- [ ] Keyboard accessible
- [ ] WCAG AA contrast
- [ ] Smooth animations

---

## 🚀 Implementation

```bash
# TODO: CLAUDE_CODE - Setup React + Tailwind

# Install dependencies
npm install tailwindcss @tailwindcss/forms
npm install lucide-react
npm install clsx tailwind-merge

# Configure Tailwind (see config above)

# Start building components
```

---

**Design System:** 🍓 Strawberry Theme  
**Status:** Ready to implement  
**Maintained by:** Ashok
