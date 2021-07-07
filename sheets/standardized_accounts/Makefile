# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make standardized_accounts.pdf

standardized_accounts.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > standardized_accounts.tex

standardized_accounts.pdf: ../../*.tex ../../trademark.pdf *.tex standardized_accounts.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode standardized_accounts.tex
	../../latexrun standardized_accounts.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: standardized_accounts.pdf
	open standardized_accounts.pdf

o:
	make open

terms: standardized_accounts.tex
	bbk_sheet -mode ts standardized_accounts.tex

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
	rm -f standardized_accounts.*
	make reset
	bbk_sheet -mode mk > Makefile