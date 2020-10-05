﻿using System;
using System.Collections.Generic;
using System.Linq;

namespace tree_level_sum
{
    class Program
    {
        static void Main(string[] args)
        {
            var root = new Tree(1);
            root.Left = new Tree(2);
            root.Right = new Tree(3);
            root.Left.Left = new Tree(4);
            root.Left.Right = new Tree(5);
            root.Right.Right = new Tree(8);
            root.Right.Right.Left = new Tree(6);
            root.Right.Right.Right = new Tree(7);
            //Breath Frist Search
            System.Console.WriteLine("================= Breath First Search on Tree ======================");
            System.Console.WriteLine($"Iterating Tree: {string.Join(',', BFS(root))}");
            System.Console.WriteLine($"Tree Level Sum: {string.Join(',', TreeLevelSum(root))}");
            System.Console.WriteLine($"Max Level Sum: {MaxLevelSum(root)}");


            //Depth First Search
            System.Console.WriteLine("================= Depth First Search on 2D Matrix ======================");
            int[][] matrix = new int[5][]
            {
                new int[5] { 1, 1, 0, 0, 0 },
                new int[5] { 0, 1, 1, 0, 0 },
                new int[5] { 0, 0, 1, 0, 1 },
                new int[5] { 1, 0, 0, 0, 1 },
                new int[5] { 0, 1, 0, 1, 1 }
            };
            System.Console.WriteLine("Print 2D Matrix!!");
            PrintMatrix(matrix);
            int r = connectedCell(matrix);
            //expect 5
            System.Console.WriteLine($"Connected cell: {r}");
        }
        static void PrintMatrix(int[][] m)
        {
            for (int i = 0; i < m.Length; i++)
            {
                for (int j = 0; j < m[i].Length; j++)
                {
                    Console.Write($"{m[i][j]} ");
                }
                System.Console.WriteLine();
            }
        }

        static int connectedCell(int[][] m)
        {
            var maxRegion = 0;
            for (int r = 0; r < m.Length; r++)
            {
                for (int c = 0; c < m[r].Length; c++)
                {
                    if (m[r][c] == 0) continue;
                    var region = getMaxRegion(m, r, c);
                    if (region > maxRegion)
                    {
                        maxRegion = region;
                    }
                }
            }
            return maxRegion;
        }
        static int getMaxRegion(int[][] m, int r, int c)
        {
            if (r < 0 || c < 0 || r >= m.Length || c >= m[r].Length)
            {
                return 0;
            }
            if (m[r][c] == 0) return 0;

            var size = 1;
            m[r][c] = 0;//mark as visited.
            for (int i = r - 1; i <= r + 1; i++)
            {
                for (int j = c - 1; j <= c + 1; j++)
                {
                    if (i == r && j == c) continue;
                    size += getMaxRegion(m, i, j);
                }
            }
            return size;
        }
        static List<int> BFS(Tree tree)
        {
            var q = new Queue<Tree>();
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

        static List<int> TreeLevelSum(Tree node)
        {
            var result = new List<int>();
            var levelNodes = new List<Tree>() { node };

            while (levelNodes.Any())
            {
                var sum = 0;
                var childNodes = new List<Tree>();
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

        static int MaxLevelSum(Tree node)
        {
            var max = 0;
            var levelNodes = new List<Tree>() { node };

            while (levelNodes.Any())
            {
                var sum = 0;
                var childNodes = new List<Tree>();
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

    public class Tree
    {
        public int Value { get; set; }
        public Tree Left { get; set; }
        public Tree Right { get; set; }
        public Tree(int value)
        {
            Value = value;
            Left = null;
            Right = null;
        }
    }

    public class Node
    {
        public string Name { get; set; }
        public List<Node> Child { get; set; }
        public Node(string name = "")
        {
            Name = name;
            Child = new List<Node>();
        }

        public Node AddChild(Node child)
        {
            Child.Add(child);
            return this;
        }
    }
}
