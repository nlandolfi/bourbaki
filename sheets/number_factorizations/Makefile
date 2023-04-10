# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make number_factorizations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

number_factorizations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > number_factorizations.tex

number_factorizations.pdf: ../../*.tex ../../trademark.pdf *.tex number_factorizations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode number_factorizations.tex
	../../latexrun number_factorizations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: number_factorizations.pdf
	open number_factorizations.pdf

o:
	make open

terms: number_factorizations.tex
	bbk sheet -mode ts number_factorizations.tex

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
	rm -f number_factorizations.*
	make reset
	bbk sheet -mode mk > Makefile