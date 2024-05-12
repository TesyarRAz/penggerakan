import Image from "next/image";

const data = [
  {
    name: "Modul 1",
    description: "Modul 1",
    submodules: [
      {
        name: "Submodul 1",
        description: "Submodul 1",
        contents: [
          {
            name: "Content 1.1",
            type: "menu",
            contents: [
              {
                name: "Blog 1",
                type: "blog",
                action: "https://google.com"
              },
              {
                name: "Blog 2",
                type: "blog",
                action: "https://google.com"
              },
              {
                name: "Forum",
                type: "forum",
                action: "https://google.com"
              }
            ]
          }
        ]
      },
      {
        name: "Submodul 2",
        description: "Submodul 2",
      },
    ]
  }
]

export default function Home() {
  return (
    <main className="flex">
      
    </main>
  );
}
