# gardi

Configurable easy data store for your things.

(Current focus prioritizes character management in multiple TTRPG systems)

(Everything is a work in progress pipe project. The goal is to get the project
done first)

## Explanation

Finding products to store your data is hard.

Dozens upon dozens of notetaking software with individual subscriptions
and quirks are found everywhere. Not to mention searching for something that
needs specific features, such as recipes, todos, movie lists, managing stats
for your favorite board game. The list goes on!

The goal of "gardi" (i.e. "to keep" in Esperanto) is not to be the perfect
system (perfect doesn't exist) but is meant to be highly configurable and
work for YOUR use case.

It works in three parts:

1. You design a "schematic," which defines what values you want to use as well
    as what can be shown and used later.
2. Configure the way the UI looks (grid, list, etc.) and associate values.
3. Use your schema! Set up your values and watch as what you set up works! Users
    can use schematics more than once and share them with others!

Rinse and repeat.

Further customization (recommended for advanced users with programming 
experience) are also in mind.

## Design

Assignment requirements for my class start here.

Nothing is set in stone, this is currently defining my "pipe dream."

### Considerations (ideas to consider)

- Frontend/Client
    - Astro: Works with a framework of choice and is meant to build sites.
        Known for its island feature, allowing selective hydration.
    - Svelte: Minimal, readable syntax that compiles to blocks of javascript.
        Works for the scale of this project and allows for lower level edits,
        which may be better for the 
    - React: More heavy, but component centered features with near unlimited
        customization and component options to choose from.
    - jsep: javascript module that parses expressions. Could be used to render
        live data on the frontend.
- Backend Tools
    - Golang: Easy to develop and debug, though initial development time is
        typically longer. Focus on code readabilty and libraries such as
        encoder/json and expr making specific tasks easier.
    - Javascript/Typescript: Very easy to get a project started. Has most 
        integrations with modern web frameworks.
    - Pocketbase: Small, minimal, fits the scope of the current state of the
        project. Expandable and easy to use database option with semi-advanced
        features.
    - Appwrite: Heavyweight but extendable suite of tools that can be self
        hosted. Includes database, auth, etc.
    - Consider front end libraries (use Next.js or SvelteKit) for this project.

### Advanced features

Basic features are specified in explanation.

- More variable types and display types
- Compatibility with REST APIs (sometimes you don't want to manually write data)
- User authentification.
- Live updates for people using the same instance.
- Schematic/instance sharing.

## Timeline

1. Nov. 8
    - Schema design and stack declaration.
    - A user should be able to conceptualize what the app is and how it works.
2. Nov. 12
    - Backend library for processing.
    - A user can start seeing an implementation being designed.
3. Nov. 15
    - API (likely done with golang/auth supporting features such as
        Pocketbase/Appwrite). Data should start getting persistently stored.
    - A user could send REST API queries and experience data persistence.
4. Nov. 19
    - Frontend design (hopefully with a framework of some kind)
    - A user should be able to build and edit schemas.
5. Nov. 22
    - Design/Polish. Tighter integration between parts and major usability bugs
        worked out between 
    - The user can not only use the application but enjoy it, with
        intuitive design and styling at the forefront.
6. Nov. 25
    - Catchup/Extra Credit/Additional Features
    - The user's mind is blown.
7. Dec. 3
    - Catchup/Extra Credit/Additional Features
    - The user's mind is blown.

## Ideas (stored for later)

Given permission to give list of ideas for now.

- Simple volleyball stat tracker
    - this would be under request from somebody
- Artist portfolio/commission site
    - Includes CDN powered image viewing
    - Form submission (for commissions/messaging)
    - Messages for commission details
    - Custom requests page
    - Payment portal
- Customizable data store (with an emphasis on TTRPG's)
    - Backend for dealing with configuration and standardization
    - Frontend for building data stores
    - Frontend for utilizing data stores
    - Interoperability for live updates (web sockets?, etc.)
    - Created for a need to work with different RPGs
- Family History UI (emphasis on self hosting)
    - Backend that supports configuration and standardization.
    - Frontend to interact with people, reference types (photo, document, etc.)
    - Ways to display reference types.
    - Log in and presentation utilization.
