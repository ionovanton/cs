#!bin/bash

job ()
{
	count=$(cat ${1} | awk 'FNR==2{ print $2 }')
	total=$(cat ${1} | awk 'FNR==2{ print $4 }')
	percent=$(awk "BEGIN { pc=100*${count}/${total}; i=int(pc); print (pc-i<0.5)?i:i+1 }")
	echo $count $total $percent
}

tt=0
cc=0
pp=0
template="https://progress-bar.dev"
r="README.md"
declare -a files=()
declare -a names=($(ls | grep -Ev '.sh|.md'))
for val in ${names[@]}; do
  files+=("${val}/${r}")
done

for s in ${files[@]}; do
	name=${s%/*}
	value=$(job $s $name)
	count=$(echo ${value} | awk '{ print $1 }')
	total=$(echo ${value} | awk '{ print $2 }')
	tt=$(( tt + total ))
	cc=$(( cc + count ))
done


echo "## Progress &nbsp;&nbsp;![Progress](${template}/${pp}/?title=${cc}/${tt})"
for s in ${files[@]}; do
	name=${s%/*}
	value=$(job $s $name)
	count=$(echo ${value} | awk '{ print $1 }')
	total=$(echo ${value} | awk '{ print $2 }')
	percent=$(echo ${value} | awk '{ print $3 }')
	echo "### ${name}"
	echo "![Progress](${template}/${percent}/?title=${count}/${total})"
done
