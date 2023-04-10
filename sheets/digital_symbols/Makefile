# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make digital_symbols.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

digital_symbols.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > digital_symbols.tex

digital_symbols.pdf: ../../*.tex ../../trademark.pdf *.tex digital_symbols.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode digital_symbols.tex
	../../latexrun digital_symbols.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: digital_symbols.pdf
	open digital_symbols.pdf

o:
	make open

terms: digital_symbols.tex
	bbk_sheet -mode ts digital_symbols.tex

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
	rm -f digital_symbols.*
	make reset
	bbk_sheet -mode mk > Makefile