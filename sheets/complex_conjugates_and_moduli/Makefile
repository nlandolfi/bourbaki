# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_conjugates_and_moduli.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

complex_conjugates_and_moduli.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > complex_conjugates_and_moduli.tex

complex_conjugates_and_moduli.pdf: ../../*.tex ../../trademark.pdf *.tex complex_conjugates_and_moduli.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_conjugates_and_moduli.tex
	../../latexrun complex_conjugates_and_moduli.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_conjugates_and_moduli.pdf
	open complex_conjugates_and_moduli.pdf

o:
	make open

terms: complex_conjugates_and_moduli.tex
	bbk sheet -mode ts complex_conjugates_and_moduli.tex

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
	bbk sheet -mode mk > Makefile