import Link from "next/link";

const Home = () => {
  return (
    <div>
      <h1>Keyboard Warrior</h1>
      <Link href="/type">
        <a>Type!</a>
      </Link>
    </div>
  );
};

export default Home;
