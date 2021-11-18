# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_conjugates_and_moduli.pdf

complex_conjugates_and_moduli.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_conjugates_and_moduli.tex

complex_conjugates_and_moduli.pdf: ../../*.tex ../../trademark.pdf *.tex complex_conjugates_and_moduli.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_conjugates_and_moduli.tex
	../../latexrun complex_conjugates_and_moduli.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_conjugates_and_moduli.pdf
	open complex_conjugates_and_moduli.pdf

o:
	make open

terms: complex_conjugates_and_moduli.tex
	bbk_sheet -mode ts complex_conjugates_and_moduli.tex

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
	rm -f complex_conjugates_and_moduli.*
	make reset
	bbk_sheet -mode mk > Makefile