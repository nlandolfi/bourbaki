# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make k-trees.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

k-trees.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > k-trees.tex

k-trees.pdf: ../../*.tex ../../trademark.pdf *.tex k-trees.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode k-trees.tex
	../../latexrun k-trees.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: k-trees.pdf
	open k-trees.pdf

o:
	make open

terms: k-trees.tex
	bbk sheet -mode ts k-trees.tex

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
	rm -f k-trees.*
	make reset
	bbk sheet -mode mk > Makefile