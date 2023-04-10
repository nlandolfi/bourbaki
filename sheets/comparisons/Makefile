# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make comparisons.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

comparisons.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > comparisons.tex

comparisons.pdf: ../../*.tex ../../trademark.pdf *.tex comparisons.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode comparisons.tex
	../../latexrun comparisons.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: comparisons.pdf
	open comparisons.pdf

o:
	make open

terms: comparisons.tex
	bbk sheet -mode ts comparisons.tex

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
	rm -f comparisons.*
	make reset
	bbk sheet -mode mk > Makefile