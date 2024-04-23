import Fuse from "fuse.js";

const list = [
  {
    title: "Old Man's War",
    author: {
      firstName: "John",
      lastName: "Scalzi",
    },
  },
  {
    title: "The Lock Artist",
    author: {
      firstName: "Steve",
      lastName: "Hamilton",
    },
  },
  {
    title: "HTML5",
    author: {
      firstName: "Remy",
      lastName: "Sharp",
    },
  },
  {
    title: "Right Ho Jeeves",
    author: {
      firstName: "P.D",
      lastName: "Woodhouse",
    },
  },
  {
    title: "The Code of the Wooster",
    author: {
      firstName: "P.D",
      lastName: "Woodhouse",
    },
  },
  {
    title: "Thank You Jeeves",
    author: {
      firstName: "P.D",
      lastName: "Woodhouse",
    },
  },
  {
    title: "The DaVinci Code",
    author: {
      firstName: "Dan",
      lastName: "Brown",
    },
  },
  {
    title: "Angels & Demons",
    author: {
      firstName: "Dan",
      lastName: "Brown",
    },
  },
  {
    title: "The Silmarillion",
    author: {
      firstName: "J.R.R",
      lastName: "Tolkien",
    },
  },
  {
    title: "Syrup",
    author: {
      firstName: "Max",
      lastName: "Barry",
    },
  },
  {
    title: "The Lost Symbol",
    author: {
      firstName: "Dan",
      lastName: "Brown",
    },
  },
  {
    title: "The Book of Lies",
    author: {
      firstName: "Brad",
      lastName: "Meltzer",
    },
  },
  {
    title: "Lamb",
    author: {
      firstName: "Christopher",
      lastName: "Moore",
    },
  },
  {
    title: "Fool",
    author: {
      firstName: "Christopher",
      lastName: "Moore",
    },
  },
  {
    title: "Incompetence",
    author: {
      firstName: "Rob",
      lastName: "Grant",
    },
  },
  {
    title: "Fat",
    author: {
      firstName: "Rob",
      lastName: "Grant",
    },
  },
  {
    title: "Colony",
    author: {
      firstName: "Rob",
      lastName: "Grant",
    },
  },
  {
    title: "Backwards, Red Dwarf",
    author: {
      firstName: "Rob",
      lastName: "Grant",
    },
  },
  {
    title: "The Grand Design",
    author: {
      firstName: "Stephen",
      lastName: "Hawking",
    },
  },
  {
    title: "The Book of Samson",
    author: {
      firstName: "David",
      lastName: "Maine",
    },
  },
  {
    title: "The Preservationist",
    author: {
      firstName: "David",
      lastName: "Maine",
    },
  },
  {
    title: "Fallen",
    author: {
      firstName: "David",
      lastName: "Maine",
    },
  },
  {
    title: "Monster 1959",
    author: {
      firstName: "David",
      lastName: "Maine",
    },
  },
];

const fuseOptions = {
  //   isCaseSensitive: false,
  //   includeScore: true,
  shouldSort: true,
  //   includeMatches: true,
  // findAllMatches: false,
  // minMatchCharLength: 0,
  // location: 0,
  // threshold: 0.0,
  // distance: 100,
  // useExtendedSearch: false,
  // ignoreLocation: false,
  // ignoreFieldNorm: false,
  // fieldNormWeight: 1,
  keys: ["title", "author.firstName"],
};

const fuse = new Fuse(list, fuseOptions);

const main = async (searchPattern: string) => {
  const res = fuse.search(searchPattern);

  console.log(JSON.stringify(res, null, 2));
};

main("artist");
