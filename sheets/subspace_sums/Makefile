# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspace_sums.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

subspace_sums.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > subspace_sums.tex

subspace_sums.pdf: ../../*.tex ../../trademark.pdf *.tex subspace_sums.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subspace_sums.tex
	../../latexrun subspace_sums.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspace_sums.pdf
	open subspace_sums.pdf

o:
	make open

terms: subspace_sums.tex
	bbk sheet -mode ts subspace_sums.tex

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
	rm -f subspace_sums.*
	make reset
	bbk sheet -mode mk > Makefile