# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_cross_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

differential_cross_entropy.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > differential_cross_entropy.tex

differential_cross_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_cross_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_cross_entropy.tex
	../../latexrun differential_cross_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_cross_entropy.pdf
	open differential_cross_entropy.pdf

o:
	make open

terms: differential_cross_entropy.tex
	bbk sheet -mode ts differential_cross_entropy.tex

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
	rm -f differential_cross_entropy.*
	make reset
	bbk sheet -mode mk > Makefile