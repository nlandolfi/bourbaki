# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make change-making_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

change-making_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > change-making_problems.tex

change-making_problems.pdf: ../../*.tex ../../trademark.pdf *.tex change-making_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode change-making_problems.tex
	../../latexrun change-making_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: change-making_problems.pdf
	open change-making_problems.pdf

o:
	make open

terms: change-making_problems.tex
	bbk sheet -mode ts change-making_problems.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f change-making_problems.*
	make reset
	bbk sheet -mode mk > Makefile