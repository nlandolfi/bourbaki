# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_sums.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_sums.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > natural_sums.tex

natural_sums.pdf: ../../*.tex ../../trademark.pdf *.tex natural_sums.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_sums.tex
	../../latexrun natural_sums.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_sums.pdf
	open natural_sums.pdf

o:
	make open

terms: natural_sums.tex
	bbk sheet -mode ts natural_sums.tex

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
	rm -f natural_sums.*
	make reset
	bbk sheet -mode mk > Makefile