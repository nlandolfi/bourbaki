# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make hidden_memory_chains.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

hidden_memory_chains.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > hidden_memory_chains.tex

hidden_memory_chains.pdf: ../../*.tex ../../trademark.pdf *.tex hidden_memory_chains.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode hidden_memory_chains.tex
	../../latexrun hidden_memory_chains.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: hidden_memory_chains.pdf
	open hidden_memory_chains.pdf

o:
	make open

terms: hidden_memory_chains.tex
	bbk_sheet -mode ts hidden_memory_chains.tex

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
	rm -f hidden_memory_chains.*
	make reset
	bbk_sheet -mode mk > Makefile