#include<stdio.h>
#include<stdlib.h>

typedef struct AVLT {
	struct AVLT* right;
	struct AVLT* left;
	int data;
	int height;
}AVL;
int getHeight(AVL* t)
{
	if (t == NULL)
		return 0;
	else
		return t->height;
}
AVL* LeftRotation(AVL* t)
{
	AVL* temp = t->right;
	t->right = t->right->left;
	temp->left = t;
	t->height = getHeight(t->left) > getHeight(t->right)
		? getHeight(t->left) + 1 :
		getHeight(t->right) + 1;
	temp->height = getHeight(temp->left) > getHeight(temp->right)
		? getHeight(temp->left) + 1 :
		getHeight(temp->right) + 1;
	return temp;
}

AVL* RightRotation(AVL* t)
{
	AVL* temp = t->left;
	t->left = t->left->right;
	temp->right = t;
	t->height = getHeight(t->left) < getHeight(t->right)
		? getHeight(t->left) + 1 :
		getHeight(t->right) + 1;
	temp->height = getHeight(temp->left) < getHeight(temp->right)
		? getHeight(temp->left) + 1 :
		getHeight(temp->right) + 1;
	return temp;
}

AVL* keepBalance(AVL* t)
{
	if (getHeight(t) <= 2)
		return t;
	if (getHeight(t->left) >= getHeight(t->right) + 2
		&& getHeight(t->left->left) > getHeight(t->left->right))
		return RightRotation(t);//需要左旋的情况
	else if (getHeight(t->right) >= getHeight(t->left) + 2
		&& getHeight(t->right->right) > getHeight(t->right->left))
		return LeftRotation(t);//需要右旋的情况
	else if (getHeight(t->left) >= getHeight(t->right) + 2
		&& getHeight(t->left->left) < getHeight(t->left->right))
	{
		t->left = LeftRotation(t->left);
		return RightRotation(t);
	}
	else if (getHeight(t->right) >= getHeight(t->left) + 2
		&& getHeight(t->right->left) < getHeight(t->right->left))
	{
		t->right = RightRotation(t->right);
		return LeftRotation(t);
	}
	else
		return t;
}
AVL* AddAVLT(int num, AVL* t)
{
	if (t == NULL)
	{
        AVL *newNode = (AVL*)malloc(sizeof(AVL));
		newNode->data = num;
        newNode->height = 1;
        newNode->left = NULL;
        newNode->right = NULL;
		return newNode;
	}
	if (num > t->data)
		t->right = AddAVLT(num, t->right);
	else
		t->left = AddAVLT(num, t->left);
	t->height = getHeight(t->left) > getHeight(t->right) ?
		getHeight(t->left) + 1:
		getHeight(t->right) + 1;
	return keepBalance(t);
}

int main()
{
	AVL T = { NULL, NULL, 13, 1 };
	for (int i = 0; i < 100000; i++)
	{
		int temp = i * (30 - i) * (i - i * (17 - i));
        // printf("%d",i);
		AddAVLT(temp, &T);
	}
    printf("%d\n", T.height);
}