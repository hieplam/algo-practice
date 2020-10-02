using System;
using System.Collections.Generic;
using System.Linq;

namespace tree_level_sum
{
    class Program
    {
        static void Main(string[] args)
        {
            var root = new Node(1);
            root.Left = new Node(2);
            root.Right = new Node(3);
            root.Left.Left = new Node(4);
            root.Left.Right = new Node(5);
            root.Right.Right = new Node(8);
            root.Right.Right.Left = new Node(6);
            root.Right.Right.Right = new Node(7);

            System.Console.WriteLine($"Iterating Tree: {string.Join(',', BFS(root))}");
            System.Console.WriteLine($"Tree Level Sum: {string.Join(',', TreeLevelSum(root))}");
            System.Console.WriteLine($"Max Level Sum: {MaxLevelSum(root)}");

        }

        static List<int> BFS(Node tree)
        {
            var q = new Queue<Node>();
            q.Enqueue(tree);
            var result = new List<int>();
            while (q.Any())
            {
                var n = q.Dequeue();
                if (n.Left != null) q.Enqueue(n.Left);
                if (n.Right != null) q.Enqueue(n.Right);

                result.Add(n.Value);
            }

            return result;
        }

        static List<int> TreeLevelSum(Node node)
        {
            var result = new List<int>();
            var levelNodes = new List<Node>() { node };

            while (levelNodes.Any())
            {
                var sum = 0;
                var childNodes = new List<Node>();
                foreach (var item in levelNodes)
                {
                    sum += item.Value;
                    if (item.Left != null) childNodes.Add(item.Left);
                    if (item.Right != null) childNodes.Add(item.Right);
                }

                result.Add(sum);
                levelNodes = childNodes;
            }

            return result;
        }

        static int MaxLevelSum(Node node){
            var max = 0;
            var levelNodes = new List<Node>() { node };

            while (levelNodes.Any())
            {
                var sum = 0;
                var childNodes = new List<Node>();
                foreach (var item in levelNodes)
                {
                    sum += item.Value;
                    if (item.Left != null) childNodes.Add(item.Left);
                    if (item.Right != null) childNodes.Add(item.Right);
                }

                max = Math.Max(max, sum);
                levelNodes = childNodes;
            }

            return max;
        }
    }

    public class Node
    {
        public int Value { get; set; }
        public Node Left { get; set; }
        public Node Right { get; set; }
        public Node(int value)
        {
            Value = value;
            Left = null;
            Right = null;
        }


    }
}
