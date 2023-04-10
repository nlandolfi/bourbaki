# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make american_standard_code_for_information_exchange.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

american_standard_code_for_information_exchange.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > american_standard_code_for_information_exchange.tex

american_standard_code_for_information_exchange.pdf: ../../*.tex ../../trademark.pdf *.tex american_standard_code_for_information_exchange.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode american_standard_code_for_information_exchange.tex
	../../latexrun american_standard_code_for_information_exchange.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: american_standard_code_for_information_exchange.pdf
	open american_standard_code_for_information_exchange.pdf

o:
	make open

terms: american_standard_code_for_information_exchange.tex
	bbk sheet -mode ts american_standard_code_for_information_exchange.tex

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
	rm -f american_standard_code_for_information_exchange.*
	make reset
	bbk sheet -mode mk > Makefile