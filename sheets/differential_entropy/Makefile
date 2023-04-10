# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

differential_entropy.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > differential_entropy.tex

differential_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_entropy.tex
	../../latexrun differential_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_entropy.pdf
	open differential_entropy.pdf

o:
	make open

terms: differential_entropy.tex
	bbk sheet -mode ts differential_entropy.tex

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
	rm -f differential_entropy.*
	make reset
	bbk sheet -mode mk > Makefile