# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make lower_upper_triangular_factorizations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

lower_upper_triangular_factorizations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > lower_upper_triangular_factorizations.tex

lower_upper_triangular_factorizations.pdf: ../../*.tex ../../trademark.pdf *.tex lower_upper_triangular_factorizations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode lower_upper_triangular_factorizations.tex
	../../latexrun lower_upper_triangular_factorizations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: lower_upper_triangular_factorizations.pdf
	open lower_upper_triangular_factorizations.pdf

o:
	make open

terms: lower_upper_triangular_factorizations.tex
	bbk sheet -mode ts lower_upper_triangular_factorizations.tex

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
	rm -f lower_upper_triangular_factorizations.*
	make reset
	bbk sheet -mode mk > Makefile